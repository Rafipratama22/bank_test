package main

import (
	"log"
	"os"

	"github.com/Rafipratama22/bank_test/router"
	"github.com/joho/godotenv"
)


func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func main() {
	godotenv.Load()
	port_server := os.Getenv("PORT_SERVER")
	router := router.Start()
	router.Run(":" + port_server)
}