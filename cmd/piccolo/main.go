package main

import (
	"log"
	"piccolo/api"

	"github.com/joho/godotenv"
)

func bootstrap() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

func main() {
	bootstrap()
	api.Start()
}
