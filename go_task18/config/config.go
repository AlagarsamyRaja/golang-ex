package config

import (
	"fmt"
	"log"
	"os"
	"sqlConnection/models"

	"github.com/joho/godotenv"
)

func Configuration() string {

	var data models.Datas

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error : Loading env file")
	}

	data.User = os.Getenv("DB_USER")
	data.Dbname = os.Getenv("DB_DATABASE")
	data.Password = os.Getenv("DB_PASSWORD")
	data.Mode = os.Getenv("PGSSLMODE")

	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", data.User, data.Dbname, data.Password, data.Mode)
	return connectionString
}
