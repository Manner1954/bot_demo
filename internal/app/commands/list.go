package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outMessageText := "Here all products: \n\n"
	products := c.productService.List()
	for _, p := range products {
		outMessageText += p.Title
		outMessageText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outMessageText)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["list"] = (*Commander).List
}
