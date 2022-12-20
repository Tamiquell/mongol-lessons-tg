package tg

import (
	"bufio"
	"context"
	"github/Tamiquell/mongol-lessons-tg/internal/messages"
	vb "github/Tamiquell/mongol-lessons-tg/internal/verbs"
	"log"
	"os"

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
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()

}

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel, msgModel *messages.Model) {
	for {
		select {
		case <-ctx.Done():
			return
		case update := <-updates:
			handleUpdate(ctx, update, msgModel)
		}
	}
}

func handleUpdate(ctx context.Context, update tgbotapi.Update, msgModel *messages.Model) {
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
