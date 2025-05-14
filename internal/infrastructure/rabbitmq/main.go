package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func ConnectToRabbitMQ(amqpURL string) *amqp.Connection {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func OpenChannelAndDeclareQueueRabbit(connection *amqp.Connection, queueName string) *amqp.Channel {
	ch, err := connection.Channel()
	if err != nil {
		log.Fatal("failed to open a channel: %w", err)
	}
	_, err = ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}
	return ch
}
