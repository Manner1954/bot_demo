package commands

import (
	"log"

	"github.com/Manner1954/bot/internal/app/commands/subdomain"
	"github.com/Manner1954/bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommanderInterface interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath)
}

type Commander struct {
	bot                *tgbotapi.BotAPI
	subdomainCommander CommanderInterface
}

func NewCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot:                bot,
		subdomainCommander: subdomain.NewSubdomainCommander(bot),
	}
}

func (c *Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "demo":
		c.subdomainCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("SubdomainCommander.HandleCallback: unknow subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *Commander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "demo":
		c.subdomainCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("Subdomain.HandleCommand: unknow command - %s", commandPath.Subdomain)
	}
}
