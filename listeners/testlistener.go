package listeners

import (
	"github.com/lordralex/mightyena/events"
	"github.com/lordralex/mightyena/logging"
	"github.com/lordralex/mightyena/services"
)

var testLogger = logging.GetLogger("TEST")

func RegisterTest() {
	services.RegisterMessage("test", messageEvent)
}

func messageEvent(event events.Message) {
	testLogger.Debug("RECEIVED: %+v", event)
	testLogger.Debug("FROM: %s", event.User().Nickname())
	if event.Channel() != nil {
		testLogger.Debug("CHANNEL: %s", event.Channel().Name())
	} else {
		testLogger.Debug("CHANNEL: NONE")
	}
	testLogger.Debug("MESSAGE: %s", event.Message())
}