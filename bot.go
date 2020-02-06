package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	storage "re9ulus.com/demetra-bot/v2/storage"
)

const (
	HELP   = "help"
	ADD    = "add"
	SPENT  = "spent"
	BUDGET = "budget"
)

func RunBot() {
	isDebug := true
	token := os.Getenv("DEMETRA_TOKEN")
	storageType, ok := os.LookupEnv("DEMETRA_STORAGE")
	if !ok {
		storageType = "memory"
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("No token provided: ")
		log.Panic(err)
	}
	bot.Debug = isDebug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	var gringotts storage.Storage
	switch storageType {
	case "memory":
		gringotts = storage.NewInMemoryStorage()
	case "redis":
		gringotts = storage.NewRedisStorage(
			redis.NewClient(
				&redis.Options{Addr: "localhost:6379"},
			),
		)
	default:
		panic("wrong storage type")
	}

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		fmt.Println(update.Message.From)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case HELP:
				msg.Text = `
                    /budget - get current budget
                    /add 42 - add amount
                    /spent 13 - spent amount
                    /help - print this message`
			case ADD:
				data, err := ParseUpdateAmountMsg(update.Message.Text)
				if err != nil {
					msg.Text = "Usage: `/add 42` Use positive integer."
					break
				}
				user_id := update.Message.From.ID
				gringotts.Add(user_id, data.val)
				msg.Text = fmt.Sprintf(
					"You added : %v, budget: %v", data.val, gringotts.Get(user_id))
			case SPENT:
				data, err := ParseUpdateAmountMsg(update.Message.Text)
				if err != nil {
					msg.Text = "Usage: `/spent 42` Use positive interger."
					break
				}
				user_id := update.Message.From.ID
				gringotts.Spent(user_id, data.val)
				msg.Text = fmt.Sprintf(
					"You spent : %v, budget: %v", data.val, gringotts.Get(user_id))
			case BUDGET:
				user_id := update.Message.From.ID
				msg.Text = fmt.Sprintf("Current budget : %v", gringotts.Get(user_id))
			default:
				msg.Text = "Wrooong! Use /help"
			}
		} else {
			msg.ReplyToMessageID = update.Message.MessageID
			msg.Text = update.Message.Text
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
