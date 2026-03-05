package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func ExtractKey() (string,error){

	err := godotenv.Load()
	if err != nil {
		return "",fmt.Errorf("Error : Loading env file : %s",err)
	}

	keys:=os.Getenv("KEY")

	return keys,nil
}