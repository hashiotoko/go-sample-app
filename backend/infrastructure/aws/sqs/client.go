package sqs

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	appCfg "github.com/hashiotoko/go-sample-app/backend/config"
	interfaces "github.com/hashiotoko/go-sample-app/backend/interfaces/clients"
)

// AWS SQSの仕様として1回のリクエストで送信できる最大メッセージ数は10
// ref. https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_SendMessageBatch.html
const MaxSendMessagesBatchSize int = 10

type client struct {
	Client    *sqs.Client
	QueueURLs map[string]*string
}

func NewClient() interfaces.MessageQueueClient {
	awsUrl := appCfg.Config.AWS.URL
	region := appCfg.Config.AWS.Region

	var cfg aws.Config
	opts := []func(*sqs.Options){}
	var err error
	if awsUrl != "" {
		opts = append(opts, func(o *sqs.Options) {
			o.BaseEndpoint = aws.String(awsUrl)
		})

		cfg, err = config.LoadDefaultConfig(
			context.TODO(),
			config.WithRegion(region),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				"test",
				"test",
				"",
			)),
		)
	} else {
		cfg, err = config.LoadDefaultConfig(
			context.TODO(),
			config.WithRegion(region),
		)
	}

	if err != nil {
		slog.Error("failed to load SQS client config", "error", err)
		panic(err)
	}

	return &client{
		Client:    sqs.NewFromConfig(cfg, opts...),
		QueueURLs: make(map[string]*string),
	}
}

func (c *client) SendMessage(ctx context.Context, queueName string, message interfaces.Message) error {
	url, err := c.getQueueUrl(ctx, queueName)
	if err != nil {
		return err
	}

	_, err = c.Client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(message.Payload),
		QueueUrl:    url,
	})
	if err != nil {
		slog.ErrorContext(ctx, "failed to send message", "queueName", queueName, "message_id", message.ID, "error", err)
		return fmt.Errorf("failed to send message: %w", err)
	}
	slog.InfoContext(ctx, "successfully sent message", "queueName", queueName, "message_id", message.ID)
	return nil
}

func (c *client) SendMessages(ctx context.Context, queueName string, messages []interfaces.Message) error {
	// ここではキューの存在チェックとURLのキャッシュのみ行う
	_, err := c.getQueueUrl(ctx, queueName)
	if err != nil {
		return err
	}

	if len(messages) == 0 {
		slog.WarnContext(ctx, "no messages to send", "queueName", queueName)
		return nil
	}

	for i := 0; i < len(messages); i += MaxSendMessagesBatchSize {
		end := i + MaxSendMessagesBatchSize
		if end > len(messages) {
			end = len(messages)
		}

		err := c.sendMessageBatch(ctx, queueName, messages[i:end])
		if err != nil {
			return err
		}
	}
	slog.InfoContext(ctx, "successfully sent messages", "queueName", queueName)

	return nil
}

func (c *client) getQueueUrl(ctx context.Context, queueName string) (*string, error) {
	if url, ok := c.QueueURLs[queueName]; ok {
		return url, nil
	}

	res, err := c.Client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{QueueName: aws.String(queueName)})
	if err != nil {
		slog.ErrorContext(ctx, "failed to get queue url", "queueName", queueName, "error", err)
		return nil, fmt.Errorf("failed to get queue url(queueName: %v): %w", queueName, err)
	}

	c.QueueURLs[queueName] = res.QueueUrl
	return res.QueueUrl, nil
}

func (c *client) sendMessageBatch(ctx context.Context, queueName string, messages []interfaces.Message) error {
	url, _ := c.getQueueUrl(ctx, queueName) // エラーはこの関数を使う側でチェック済みのため無視
	entries := make([]types.SendMessageBatchRequestEntry, 0, len(messages))
	for _, entry := range messages {
		entries = append(entries, types.SendMessageBatchRequestEntry{
			Id:          aws.String(entry.ID),
			MessageBody: aws.String(entry.Payload),
		})
	}

	res, err := c.Client.SendMessageBatch(ctx, &sqs.SendMessageBatchInput{
		Entries:  entries,
		QueueUrl: url,
	})
	if err != nil {
		slog.ErrorContext(ctx, "failed to send messages", "queueName", queueName, "error", err)
		return fmt.Errorf("failed to send messages(queueName: %v): %w", queueName, err)
	}

	for _, success := range res.Successful {
		slog.InfoContext(ctx, "successfully sent message", "queueName", queueName, "message_id", aws.ToString(success.Id))
	}
	for _, failed := range res.Failed {
		slog.ErrorContext(ctx, "failed to send message", "queueName", queueName, "message_id", aws.ToString(failed.Id), "error", aws.ToString(failed.Message))
	}

	return nil
}
