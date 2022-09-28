package router

import (
	"log"
	"runtime/debug"

	"github.com/Manner1954/bot/internal/app/commands"
	"github.com/Manner1954/bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommanderInterface interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath)
}

type Router struct {
	bot                *tgbotapi.BotAPI
	subdomainCommander CommanderInterface
}

func NewRouter(bot *tgbotapi.BotAPI) *Router {
	return &Router{
		bot:                bot,
		subdomainCommander: commands.NewCommander(bot),
	}
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Recovered from panic: %v\n%v", panicValue, string(debug.Stack()))
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		r.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		r.handleMessage(update.Message)
	}
}

func (r *Router) handleCallback(callback *tgbotapi.CallbackQuery) {

	callbackPath, err := path.ParseCallback(callback.Data)

	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", callback.Data, err)
		return
	}

	switch callbackPath.Subdomain {
	case "demo":
		r.subdomainCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("Router.handleCallback: Subdomain %s not found", callbackPath.Subdomain)
	}
}

func (r *Router) handleMessage(msg *tgbotapi.Message) {
	if !msg.IsCommand() {
		r.showCommandFormat(msg)
		return
	}

	commandPath, err := path.ParseCommand(msg.Command())
	if err != nil {
		log.Printf("Router.handleMessage: error parsing message command `%s` - %v", msg.Command(), err)
		return
	}

	switch commandPath.Subdomain {
	case "demo":
		r.subdomainCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("Router.handleMessage: Subdomain %s not found", commandPath.Subdomain)
	}
}

func (r *Router) showCommandFormat(inputMsg *tgbotapi.Message) {
	outputMsg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Command format: /command_subdomain")
	_, err := r.bot.Send(outputMsg)
	if err != nil {
		log.Printf("Router.ShowCommand: error sending reply message to chat - %s", err)
	}
}
