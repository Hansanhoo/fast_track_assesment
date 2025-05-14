package rabbitmq

import (
	"asssesment_fast_track/internal/domain/models"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

func PublishPayments(ch *amqp.Channel, queueName string, data []models.Payment, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		for _, payment := range data {
			body, err := json.Marshal(payment)
			if err != nil {
				log.Printf("Failed to marshal payment: %v", err)
				continue
			}

			err = ch.Publish(
				"",        // exchange
				queueName, // key (queue name)
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType: "application/json",
					Body:        body,
				})
			if err != nil {
				log.Printf("Failed to publish message: %v", err)
			} else {
				fmt.Printf("Published: %+v\n", payment)
			}
		}
		log.Println("Payments have been published.")

	}()
	wg.Wait()
}
