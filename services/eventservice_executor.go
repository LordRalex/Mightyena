package services

import (
	"github.com/lordralex/mightyena/events"
	"github.com/lordralex/mightyena/logging"
)

var listenerMapping = make(map[string]map[string][]func(events.Event))
var eventLogger = logging.GetLogger("EVENT-SVC")

func register(eventName string, moduleName string, function func(event events.Event)) {
	if listenerMapping[eventName] == nil {
		listenerMapping[eventName] = make(map[string][]func(message events.Event))
	}

	listenerMapping[eventName][moduleName] = append(listenerMapping[eventName][moduleName], function)
}

func fireEvent(event events.Event) {
	executors := listenerMapping[event.EventName()]
	eventLogger.Debug("Firing event")
	eventLogger.Debug("%+v", event)
	for k, f := range executors {
		eventLogger.Debug("Running executors for %s", k)
		for _, function := range f {
			func() {
				defer func() {
					if err := recover(); err != nil {
						eventLogger.Error("Error running event %s: %s", event.EventName(), err)
					}
				}()
				function(event)
			}()
		}
	}
}
