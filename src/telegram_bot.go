package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"gopkg.in/telegram-bot-api.v4"
)

var bot *tgbotapi.BotAPI

func runTelegramBot(token string) {

	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("telegram bot: %s", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		message := update.Message.Text

		var currentUser User
		db.FirstOrCreate(
			&currentUser,
			User{
				Model: gorm.Model{
					ID: uint(update.Message.From.ID)},
				UserName:  update.Message.From.UserName,
				LastName:  update.Message.From.LastName,
				FirstName: update.Message.From.FirstName})

		reply := "Unknown command. Type /help"
		commands := strings.Fields(message)

		switch commands[0] {
		case "/start":
			reply = fmt.Sprintf("Type /help for help. ")

		case "/help":
			reply = fmt.Sprint("")

		case "/lock":
			reply = lockState(update.Message.From, commands[1:])

		case "/subscribtion":
			reply = subscribtionState(update.Message.From, commands[1:])
		}

		sendMessage(update.Message.Chat.ID, reply)
	}
}

func sendMessage(ID int64, message string) {
	bot.Send(tgbotapi.NewMessage(ID, message))
}

func lockState(user *tgbotapi.User, commands []string) string {
	command := commands[0]
	var reply string
	switch command {
	case "add":
		db.Create(&Lock{Serial: commands[1]})
		reply = fmt.Sprintf("Add new lock %s", commands[1])
	case "remove":
		db.Where("serial = ?", commands[1]).Delete(&Lock{})
		reply = fmt.Sprintf("Remove lock %s", commands[1])
	default:
		reply = "Unknown command"
	}
	return reply
}

func subscribtionState(user *tgbotapi.User, commands []string) string {
	return "Unknown command"
}
