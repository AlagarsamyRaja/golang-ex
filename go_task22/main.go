package main

import (
	"fmt"
	"gorm/config"
	"gorm/pkg"
	"gorm/router"
	"net/http"
)

func main() {
	connectionKey, err := config.Configuration()
	if err != nil {
		fmt.Println(err)
	}

	db, err := pkg.InitDB(connectionKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Routes(db)
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
