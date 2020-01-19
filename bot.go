package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

const (
	HELP   = "help"
	ADD    = "add"
	SPENT  = "spent"
	BUDGET = "budget"
)


func create_bot() {
	token := os.Getenv("DEMETRA_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
        log.Printf("No token provided: ")
        log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	var budget int64 = 0

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		// update.Message.Text
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
				budget += data.val
				msg.Text = fmt.Sprintf(
                    "You added : %v, budget: %v", data.val, budget)
			case SPENT:
                data, err := ParseUpdateAmountMsg(update.Message.Text)
				if err != nil {
                    msg.Text = "Usage: `/spent 42` Use positive interger."
					break
				}
                budget -= data.val
				msg.Text = fmt.Sprintf(
                    "You spent : %v, budget: %v", data.val, budget)
			case BUDGET:
				msg.Text = fmt.Sprintf("Current budget : %v", budget)
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

func main() {
	create_bot()
}
