package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func ConsumeRabbitQueue(rmqChannel *amqp.Channel, queueName string) <-chan amqp.Delivery {
	msgs, err := rmqChannel.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	fmt.Println("Waiting for messages...")

	return msgs
}
