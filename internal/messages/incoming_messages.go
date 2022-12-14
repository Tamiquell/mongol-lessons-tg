package messages

import (
	"context"
	hp "github/Tamiquell/mongol-lessons-tg/internal/helpers"
	vb "github/Tamiquell/mongol-lessons-tg/internal/verbs"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageSender interface {
	SendMessage(text string, userID int64, keyboard tgbotapi.ReplyKeyboardMarkup) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Message struct {
	Text   string
	UserID int64
}

var previousCommand string
var answers = make(map[int64]string)
var currentlyIn string

func (s *Model) IncomingMessage(ctx context.Context, msg Message) error {
	switch msg.Text {
	case "/start":
		currentlyIn = "/start"
		return s.tgClient.SendMessage(startMessage, msg.UserID, hp.StartKeyboard)
	case "/help":
		currentlyIn = "/start"
		return s.tgClient.SendMessage(helpMessage, msg.UserID, hp.StartKeyboard)
	case "/verbs":
		currentlyIn = "/verbs"
		return s.tgClient.SendMessage("You can learn new verbs here", msg.UserID, hp.VerbsInitKeyboard)
	case "/study":
		currentlyIn = "/study"
		return s.tgClient.SendMessage("Choose letters range", msg.UserID, hp.StudyKeyboard)
	case "а - б", "г - з", "и - ө", "с - т", "у - х", "ц - я":
		currentlyIn = "/study"
		row := strings.Split(msg.Text, " - ")
		first, second := row[0], row[1]
		verbsList, err := vb.VerbsList(first, second)
		if err != nil {
			return err
		}
		return s.tgClient.SendMessage(verbsList, msg.UserID, hp.StudyKeyboard)
	case "/back":
		if currentlyIn == "/verbs" {
			return s.tgClient.SendMessage("You are back in main menu", msg.UserID, hp.StartKeyboard)
		} else if currentlyIn == "/test" {
			currentlyIn = "/verbs"
			return s.tgClient.SendMessage("You are back in verbs menu", msg.UserID, hp.VerbsInitKeyboard)
		} else if currentlyIn == "/study" {
			currentlyIn = "/verbs"
			return s.tgClient.SendMessage("You are back in verbs menu", msg.UserID, hp.VerbsInitKeyboard)
		}

	case "/test":
		currentlyIn = "/test"
		return s.tgClient.SendMessage("Test your knowledge", msg.UserID, hp.VerbsTestKeyboard)
	case "/reroll":
		currentlyIn = "/test"
		text, answ, err := vb.RandomVerbs()
		answers[msg.UserID] = answ
		if err != nil {
			return err
		}
		return s.tgClient.SendMessage(text, msg.UserID, hp.VerbsTestKeyboard)
	case "/answers":
		currentlyIn = "/test"
		return s.tgClient.SendMessage(answers[msg.UserID], msg.UserID, hp.VerbsTestKeyboard)
	}

	return nil
}
