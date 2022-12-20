package tg

import (
	"bufio"
	"context"
	"github/Tamiquell/mongol-lessons-tg/internal/messages"
	vb "github/Tamiquell/mongol-lessons-tg/internal/verbs"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type Client struct {
	client *tgbotapi.BotAPI
}

func New(token string) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, errors.Wrap(err, "NewBotAPI")
	}
	return &Client{
		client: client,
	}, nil
}

func (c *Client) SendMessage(text string, userID int64, keyboard tgbotapi.ReplyKeyboardMarkup) error {
	newMessage := tgbotapi.NewMessage(userID, text)
	newMessage.ReplyMarkup = keyboard
	_, err := c.client.Send(newMessage)
	if err != nil {
		return errors.Wrap(err, "client,.end")
	}
	return nil
}

func (c *Client) ListenUpdates(msgModel *messages.Model) {

	vb.FillVerbs()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	updates := c.client.GetUpdatesChan(u)

	log.Println("listening for messages")
	go receiveUpdates(ctx, updates, msgModel)
	time.Sleep(time.Millisecond * 50)
	bufio.NewReader(os.Stdin).ReadBytes('q')
	cancel()

}

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel, msgModel *messages.Model) {
	log.Println("inside receiveUpdates")
	for {
		select {
		case update := <-updates:
			log.Println("inside update := <- updates")
			handleUpdate(ctx, update, msgModel)
		case <-ctx.Done():
			return
		}
	}
}

func handleUpdate(ctx context.Context, update tgbotapi.Update, msgModel *messages.Model) {
	log.Println("inside handleUpdate")
	if update.Message != nil { // If we got a message
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		err := msgModel.IncomingMessage(ctx, messages.Message{
			Text:   update.Message.Text,
			UserID: update.Message.From.ID,
		})
		if err != nil {
			log.Println("error processing message:", err)
		}
	}
}
