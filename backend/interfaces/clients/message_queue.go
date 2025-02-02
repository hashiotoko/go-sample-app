package interfaces

import "context"

type MessageQueueClient interface {
	SendMessage(ctx context.Context, queueName string, message Message) error
	SendMessages(ctx context.Context, queueName string, messages []Message) error
}

type Message struct {
	ID      string
	Payload string
}
