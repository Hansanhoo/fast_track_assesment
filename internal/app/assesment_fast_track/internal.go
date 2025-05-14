package logic

import (
	"asssesment_fast_track/internal/domain/models"
	mydbsql "asssesment_fast_track/internal/infrastructure/mydbsql"
	rabbitmq "asssesment_fast_track/internal/infrastructure/rabbitmq"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func StartApp() {
	// get env Variables
	mysqlUser := os.Getenv("MYSQLUSER")
	mysqlPassword := os.Getenv("MYSQLPASSWORD")
	rabbitUser := os.Getenv("RABBITUSER")
	rabbitPassword := os.Getenv("RABBITPASSWORD")

	queueName := "payments"

	connStrRabbit := fmt.Sprintf("amqp://%s:%s@rabbitmq:5672/", rabbitUser, rabbitPassword)
	connStrMYsql := fmt.Sprintf("%s:%s@tcp(mysql:3306)/mydb", mysqlUser, mysqlPassword)

	// Connect to RabbitMQ
	conn := rabbitmq.ConnectToRabbitMQ(connStrRabbit)

	defer conn.Close()

	// Create a Rabbit channel
	ch := rabbitmq.OpenChannelAndDeclareQueueRabbit(conn, queueName)

	defer ch.Close()

	var wg sync.WaitGroup
	// Publish the payments
	rabbitmq.PublishPayments(ch, queueName, models.GetMockPayments(), &wg)
	//connect to MYSQL
	mysqldb := mydbsql.ConnectMysql(connStrMYsql)

	//create payments and skippedmessagetable
	mydbsql.CreatePaymentsTable(mysqldb, &wg)
	mydbsql.CreateSkippedMessagesTable(mysqldb, &wg)

	// consume rabbit queue
	msgs := rabbitmq.ConsumeRabbitQueue(ch, queueName)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for d := range msgs {
			var p models.Payment
			if err := json.Unmarshal(d.Body, &p); err != nil {
				log.Printf("JSON decode error: %v", err)
				continue
			}

			mydbsql.InsertPayment(mysqldb, p)
		}

	}()

	go func() {
		defer wg.Done()
		//wait a bit to insert duplicate
		time.Sleep(2 * time.Second)

		mydbsql.InsertPayment(mysqldb, models.GetMockDuplicatePayment())

	}()
	wg.Wait()
}
