package main

import (
	"context"
	"github/Tamiquell/mongol-lessons-tg/internal/tg"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github/Tamiquell/mongol-lessons-tg/internal/messages"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancelFn := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	tgClient, err := tg.New(os.Getenv("TELEGRAM_APITOKEN"))
	// tgClient, err := tg.New(os.Getenv("TELEGRAM_APITOKEN_DEV"))

	if err != nil {
		log.Fatal("tg client init failed:", err)
	}
	msgModel := messages.New(tgClient)
	go tgClient.ListenUpdates(ctx, &wg, msgModel)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	cancelFn()
}
