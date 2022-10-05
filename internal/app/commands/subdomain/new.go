package subdomain

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *SubdomainCommander) New(inputMessage *tgbotapi.Message) {
	arg := inputMessage.CommandArguments()

	if arg == "" {
		log.Printf("SubdomainCommander.New: command had wrong amount of parameters")
		return
	}

	res, err := c.productService.New(arg)
	if err != nil {
		log.Printf("SubdomainCommander.New: service had error for create entity - %s", err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, res)
	_, err = c.bot.Send(msg)

	if err != nil {
		log.Printf("SubdomainCommander.New: error sending reply message to chat - %v", err)
	}
}
