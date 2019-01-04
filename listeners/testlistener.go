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
	testLogger.Info("FROM: %s", event.User().Nickname())
	if event.Channel() != nil {
		testLogger.Info("CHANNEL: %s", event.Channel().Name())
	} else {
		testLogger.Info("CHANNEL: NONE")
	}
	testLogger.Info("MESSAGE: %s", event.Message())
}