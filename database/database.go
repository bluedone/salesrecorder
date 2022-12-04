package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	config "sr-server/config"
)

// create connection with postgres db
func CreateConnection() *sql.DB {

	if config.Err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", config.PsqlInfo)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected!")
	// return the connection
	return db
}
