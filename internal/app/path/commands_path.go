package path

import (
	"errors"
	"fmt"
	"strings"
)

type CommandPath struct {
	CommandName string
	Subdomain   string
}

var errorUnknowCommand = errors.New("Unknow command")

func ParseCommand(commandText string) (CommandPath, error) {
	commandParts := strings.SplitN(commandText, "_", 2)

	if len(commandParts) != 2 {
		return CommandPath{}, errorUnknowCommand
	}

	return CommandPath{
		CommandName: commandParts[0],
		Subdomain:   commandParts[1],
	}, nil
}

func (p CommandPath) String() string {
	return fmt.Sprintf("%s_%s", p.CommandName, p.Subdomain)
}
