package sqs

//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	appCfg "github.com/hashiotoko/go-sample-app/backend/config"
)

type Client interface {
	SendMessage(ctx context.Context, queueUrl string, message string) error
	SendBatchMessages(ctx context.Context, queueUrl string, messages []string) error
}

type client struct {
	Client *sqs.Client
}

func NewClient() Client {
	region := appCfg.Config.AWS.Region
	endpoint := appCfg.Config.AWS.URL

	var cfg aws.Config
	var err error
	if endpoint != "" {
		cfg, err = config.LoadDefaultConfig(
			context.TODO(),
			config.WithRegion(region),
			config.WithBaseEndpoint(endpoint),
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
		slog.Error("failed to get SQS client config", "error", err)
		panic(err)
	}

	return client{
		Client: sqs.NewFromConfig(cfg),
	}
}

func (c client) SendMessage(ctx context.Context, queueURL string, message string) error {
	_, err := c.Client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(queueURL),
	})
	if err != nil {
		slog.Error("failed to send message", "error", err)
		return err
	}
	return nil
}

func (c client) SendBatchMessages(ctx context.Context, queueURL string, messages []string) error {
	entries := make([]types.SendMessageBatchRequestEntry, 0, len(messages))
	for i, message := range messages {
		entries = append(entries, types.SendMessageBatchRequestEntry{
			Id:          aws.String(strconv.Itoa(i)),
			MessageBody: aws.String(message),
		})
	}

	_, err := c.Client.SendMessageBatch(ctx, &sqs.SendMessageBatchInput{
		Entries:  entries,
		QueueUrl: aws.String(queueURL),
	})
	if err != nil {
		slog.Error("failed to send batch messages", "error", err)
		return err
	}
	return nil
}
