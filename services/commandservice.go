package services

import (
	"github.com/lordralex/mightyena/events"
	"strings"
)

var commands = make(map[string]*registeredCommand)

type registeredCommand struct {
	module   string
	executor func(events.Command)
}

func executeCommand(event events.Command) {
	cmd := commands[strings.ToLower(event.Command())]

	if cmd != nil {
		cmd.executor(event)
	}
}

func RegisterCommand(module, command string, function func(events.Command)) {
	commands[strings.ToLower(command)] = &registeredCommand{
		module:   module,
		executor: function,
	}
}
