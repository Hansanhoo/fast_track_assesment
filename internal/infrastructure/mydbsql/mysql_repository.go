package mydbsql

import (
	"asssesment_fast_track/internal/domain/models"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
)

func CreatePaymentsTable(db *sql.DB, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		createTableSQL := `
			CREATE TABLE IF NOT EXISTS payment_events (
				payment_id INT PRIMARY KEY,
				user_id INT NOT NULL,
				deposit_amount INT NOT NULL
			);`

		executeSql(db, createTableSQL)
	}()

	wg.Wait()
	fmt.Println("Table 'payment_events' created (if not exists).")
}

func CreateSkippedMessagesTable(db *sql.DB, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		createTableSQL := `
			CREATE TABLE IF NOT EXISTS skipped_messages (
				payment_id INT PRIMARY KEY,
				user_id INT NOT NULL,
				deposit_amount INT NOT NULL
			);`

		executeSql(db, createTableSQL)
	}()

	wg.Wait()

	fmt.Println("Table 'skipped messages' created (if not exists).")

}

func executeSql(db *sql.DB, sql string) {
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatalf("Failed to execute sql: %v", err)
	}
}

func InsertPayment(db *sql.DB, p models.Payment) error {

	_, err := db.Exec(`INSERT INTO payment_events (user_id, payment_id, deposit_amount) VALUES (?, ?, ?)`,
		p.UserID, p.PaymentID, p.DepositAmount)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			log.Printf("Primary key violation for payment_id=%d. Inserting into skipped_messages.", p.PaymentID)
			insertSkippedMessage(db, p)
		} else {
			if err := db.Ping(); err != nil {
				log.Fatalf("MySQL3 ping error: %v", err)
			}
			log.Printf("Insert error: %v", err)
		}
	} else {
		fmt.Printf("Inserted into payment_events: %+v\n", p)
	}
	return err
}

func insertSkippedMessage(db *sql.DB, p models.Payment) {
	_, err := db.Exec(`INSERT INTO skipped_messages (user_id, payment_id, deposit_amount) VALUES (?, ?, ?)`,
		p.UserID, p.PaymentID, p.DepositAmount)
	if err != nil {
		log.Printf("Failed to insert into skipped_messages: %v", err)
	} else {
		fmt.Printf("Inserted into skipped_messages: %+v\n", p)
	}
}
