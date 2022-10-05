package subdomain

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *SubdomainCommander) Remove(inputMessage *tgbotapi.Message) {
	outMessageText := "Remove product: \n\n"

	args := inputMessage.CommandArguments()
	arg, err := strconv.Atoi(args)

	if err != nil {
		outMessageText += fmt.Sprintf("Remove.Remove: strconv.Atoi returned a error - %s", err)
	}

	if err == nil {
		retString, err := c.productService.Remove(arg)

		if err != nil {
			outMessageText += fmt.Sprintf("Remove.Remove: productService.Remove returned a error - %s", err)
		} else {
			outMessageText += retString
		}
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outMessageText)

	c.bot.Send(msg)
}
