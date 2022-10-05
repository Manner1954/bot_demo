package subdomain

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *SubdomainCommander) Get(inputMessage *tgbotapi.Message) {
	outMessageText := "Get product: \n\n"

	args := inputMessage.CommandArguments()
	arg, err := strconv.Atoi(args)

	if err != nil {
		outMessageText += fmt.Sprintf("Get.Get: strconv.Atoi returned a error - %s", err)
	}

	if err == nil {
		product, err := c.productService.Get(arg)

		if err != nil {
			outMessageText += fmt.Sprintf("Get.Get: productService.Get returned a error - %s", err)
		} else {
			outMessageText += product.Title
		}
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outMessageText)

	c.bot.Send(msg)
}
