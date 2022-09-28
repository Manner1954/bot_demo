package subdomain

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Manner1954/bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *SubdomainCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)

	if err != nil {
		log.Printf("SubdomainCommnder.CallbakList: "+
			"error reading json for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)

		return
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID,
		fmt.Sprintf("Parsed: %+v\n", parsedData))

	_, err = c.bot.Send(msg)

	if err != nil {
		log.Printf("SubdomainCommnder.CallbakList: error sending reply message to chat - %v", err)
	}
}
