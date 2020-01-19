package main

import (
	"errors"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	HELP   = "help"
	ADD    = "add"
	SPENT  = "spent"
	BUDGET = "budget"
)

func validate_amount(val int64) bool {
	return val > 0
}

// TODO: Parse data to data-structure fields
func extract_amount(s string) (int64, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return 0, errors.New("Wrong string format, expected 2 tokens separated with space")
	}
	amount, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, errors.New("Can not parse amount, second token must be integer")
	}
	return amount, nil
}

func create_bot() {
	token := os.Getenv("DEMETRA_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
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
				amount, err := extract_amount(update.Message.Text)
				if err != nil {
					msg.Text = "Usage: /add 42"
					break
				}
				if !validate_amount(amount) {
					msg.Text = fmt.Sprintf("Wrong amount %v, use positive integer", amount)
					break
				}
				budget += amount
				msg.Text = fmt.Sprintf("You added : %v, budget: %v", amount, budget)
			case SPENT:
				amount, err := extract_amount(update.Message.Text)
				if err != nil {
					msg.Text = "Usage: /spent 42"
					break
				}
				if !validate_amount(amount) {
					msg.Text = fmt.Sprintf("Wrong amount %v, use positive integer")
					break
				}
				budget -= amount
				msg.Text = fmt.Sprintf("You spent : %v, budget: %v", amount, budget)
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
