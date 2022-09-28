package subdomain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *SubdomainCommander) Get(inputMessage *tgbotapi.Message) {
	outMessageText := "Get product: \n\n"
	products := c.productService.List()
	for _, p := range products {
		outMessageText += p.Title
		outMessageText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outMessageText)

	c.bot.Send(msg)
}
