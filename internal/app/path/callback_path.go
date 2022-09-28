package path

import (
	"errors"
	"fmt"
	"strings"
)

type CallbackPath struct {
	Subdomain    string
	CallbackName string
	CallbackData string
}

var errorUnknowCallback = errors.New("Unknow callback")

func ParseCallback(callbackData string) (CallbackPath, error) {
	callbackParts := strings.SplitN(callbackData, "_", 3)

	if len(callbackParts) != 3 {
		return CallbackPath{}, errorUnknowCallback
	}

	return CallbackPath{
		Subdomain:    callbackParts[0],
		CallbackName: callbackParts[1],
		CallbackData: callbackParts[2],
	}, nil
}

func (p CallbackPath) String() string {
	return fmt.Sprintf("%s_%s_%s", p.Subdomain, p.CallbackName, p.CallbackData)
}
