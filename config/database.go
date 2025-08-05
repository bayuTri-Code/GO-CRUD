package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)


var DB *sql.DB

func ConnectDb() {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database successfully")
	DB = db
}
