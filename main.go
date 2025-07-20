package main

import (
	"dottest/cmd"
	"dottest/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	log.Printf("Default config: %v", config.DefaultConfig)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cmd.Execute()
}
