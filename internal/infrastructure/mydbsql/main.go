package mydbsql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMysql(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("MySQL ping error: %v", err)
	}
	return db

}
