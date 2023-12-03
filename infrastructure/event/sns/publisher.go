package sns

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"grafana-api/infrastructure/secretmanager"
)

type IEventPublisher interface {
	SendMessage(ctx context.Context, message string) error
}

type eventPublisher struct {
	TopicName string
}

func NewEventPublisher(secret secretmanager.TopicSecret) IEventPublisher {
	return &eventPublisher{
		TopicName: secret.TopicName,
	}
}

func (e eventPublisher) SendMessage(ctx context.Context, message string) error {
	sess, err := session.NewSession()
	if err != nil {
		return fmt.Errorf("session error: %w", err)
	}

	input := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(e.TopicName),
	}

	if _, err = sns.New(sess).Publish(input); err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	return nil
}
