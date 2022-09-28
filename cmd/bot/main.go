package main

import (
	"log"
	"os"

	"github.com/Manner1954/bot/internal/app/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	token, found := os.LookupEnv("TOKEN")

	if !found {
		log.Panic("Token isn't found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	router := router.NewRouter(bot)

	for update := range updates {
		router.HandleUpdate(update)
	}
}
