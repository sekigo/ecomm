package main

import (
	"log"

	"github.com/sekigo/ecomm/db"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Successfully connected to the database")

}
