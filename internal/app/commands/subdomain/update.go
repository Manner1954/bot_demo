package subdomain

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *SubdomainCommander) Update(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	argParts := strings.SplitAfterN(args, " ", 2)

	if len(argParts) != 2 {
		log.Printf("SubdomainCommander.Update: command had wrong amount of parameters")
		return
	}

	idx, err := strconv.Atoi(strings.TrimSpace(argParts[0]))
	if err != nil {
		log.Printf("SubdomainCommander.Update: error parsing index of entity for command - %s", err)
		return
	}

	res, err := c.productService.Update(idx, argParts[1])
	if err != nil {
		log.Printf("SubdomainCommander.Update: service had error for updating entity - %s", err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, res)
	_, err = c.bot.Send(msg)

	if err != nil {
		log.Printf("SubdomainCommander.Update: error sending reply message to chat - %v", err)
	}
}
