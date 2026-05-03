package queue

import (
	"context"

	"github.com/pahan-fe/lite-streaming/backend/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	channel *amqp.Channel
}

func (r *RabbitMQ) Publish(ctx context.Context, queueName string, body []byte) error {
	_, err := r.channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = r.channel.PublishWithContext(ctx, "", queueName, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	return err
}

func (r *RabbitMQ) Consume(ctx context.Context, queueName string) (<-chan amqp.Delivery, error) {
	_, err := r.channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	messages, err := r.channel.ConsumeWithContext(ctx, queueName, "", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func NewRabbitMQ(c *config.Config) (*RabbitMQ, error) {
	conn, err := amqp.Dial(c.RabbitMQURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{channel: ch}, nil
}
