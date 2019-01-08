package services

import (
	"github.com/lordralex/mightyena/events"
)

func RegisterAction(moduleName string, function func(message events.Action)) {
	register("action", moduleName, wrapAction(function))
}

func RegisterMessage(moduleName string, function func(message events.Message)) {
	register("message", moduleName, wrapMessage(function))
}