package main

import (
	"log/slog"
	"piccolo/api"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Debug("Error loading .env file", "error", err)
	}
}

func main() {
	api.Start()
}
