package tg

import (
	"context"
	"github/Tamiquell/mongol-lessons-tg/internal/messages"
	vb "github/Tamiquell/mongol-lessons-tg/internal/verbs"
	"log"
	"sync"

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

func (c *Client) ListenUpdates(ctx context.Context, wg *sync.WaitGroup, msgModel *messages.Model) {

	wg.Add(1)
	defer wg.Done()

	vb.FillVerbs()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)

	log.Println("listening for messages")

LOOP:
	for {
		select {
		case update := <-updates:
			if update.Message != nil {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				err := msgModel.IncomingMessage(ctx, messages.Message{
					Text:   update.Message.Text,
					UserID: update.Message.From.ID,
				})

				if err != nil {
					log.Println("error processing message:", err)
				}
			}
		case <-ctx.Done():
			break LOOP
		}
	}

}
