package main

import (
	"fmt"
	"log"

	"github.com/currencies-exchange/models"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Server is up and running!")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.InitialMigration()
	router()
}
