package subdomain

import (
	"log"

	"github.com/Manner1954/bot/internal/app/path"
	"github.com/Manner1954/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SubdomainCommander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewSubdomainCommander(bot *tgbotapi.BotAPI) *SubdomainCommander {

	productService := product.NewService()
	return &SubdomainCommander{
		bot:            bot,
		productService: productService,
	}
}

func (c *SubdomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("SubdomainCommander.HandleCallback: unknow ccallback name: %s", callbackPath.CallbackName)
	}
}

func (c *SubdomainCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {

	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
