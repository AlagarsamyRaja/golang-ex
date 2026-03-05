package pkg

import (
	"database/sql"
	"log"
	"sqlConnection/config"
)

var DB *sql.DB

func InitDb() {

	connection := config.Configuration()

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatalf("database connection error %s", err)
	}

	DB = db
}
