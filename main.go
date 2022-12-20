package main

import (
	"github/Tamiquell/mongol-lessons-tg/internal/tg"
	"log"
	"os"

	"github/Tamiquell/mongol-lessons-tg/internal/messages"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// tgClient, err := tg.New(os.Getenv("TELEGRAM_APITOKEN"))
	tgClient, err := tg.New(os.Getenv("TELEGRAM_APITOKEN_DEV"))

	if err != nil {
		log.Fatal("tg client init failed:", err)
	}
	msgModel := messages.New(tgClient)
	tgClient.ListenUpdates(msgModel)
}
