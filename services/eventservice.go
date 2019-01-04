package services

import (
	"github.com/lordralex/mightyena/events"
	"github.com/lordralex/mightyena/logging"
	"github.com/thoj/go-ircevent"
)

var listenerMapping = make(map[string]map[string][]func(events.Event))
var eventLogger = logging.GetLogger("EVENT SERVICE")

func RegisterMessage(moduleName string, function func(message events.Message)) {
	register("message", moduleName, func(event events.Event) {
		function(event.(events.Message))
	})
}

func register(eventName string, moduleName string, function func(event events.Event)) {
	if listenerMapping[eventName] == nil {
		listenerMapping[eventName] = make(map[string][]func(message events.Event))
	}

	listenerMapping[eventName][moduleName] = append(listenerMapping[eventName][moduleName], function)
}

func fireMessageEvent(event *irc.Event) {
	message := event.Message()
	nick := event.Nick
	channel := event.Arguments[0]

	evt := events.CreateMessageEvent(message, GetUser(nick), GetChannel(channel), event.Connection)
	fireEvent(evt)
}

func fireEvent(event events.Event) {
	executors := listenerMapping[event.EventName()]
	eventLogger.Debug("Firing event")
	eventLogger.Debug("%+v", event)
	for k, f := range executors {
		eventLogger.Debug("Running executors for %s", k)
		for _, function := range f {
			function(event)
		}
	}
}