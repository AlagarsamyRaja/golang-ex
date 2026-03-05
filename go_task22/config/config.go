package config

import (
	"fmt"
	"gorm/models"
	"os"
	"github.com/joho/godotenv"
)

func Configuration() (string,error) {

	var data models.Datas

	err := godotenv.Load()
	if err != nil {
		return "",fmt.Errorf("Error:Loading env file")
	}

	data.Host = os.Getenv("DB_HOST")
	data.User = os.Getenv("DB_USER")
	data.Dbname = os.Getenv("DB_DATABASE")
	data.Password = os.Getenv("DB_PASSWORD")
	data.Port = os.Getenv("DB_PORT")
	data.Mode = os.Getenv("PGSSLMODE")

	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s", data.Host, data.User, data.Dbname, data.Password, data.Port, data.Mode)
	return connectionString,nil
}
