package welcome

import (
	"github.com/lordralex/mightyena/events"
	"github.com/lordralex/mightyena/services"
	"strings"
)

const ModuleName = "welcome"

var mapping = make(map[string]string)

func Load() {
	//services.RegisterCommand(ModuleName, "welcome", runCommand)
	//services.RegisterJoin(ModuleName, runJoin)
}

func runCommand(event *events.Command) {
	if event.User.NickservAccount() == "" {
		return
	}

	if len(event.Arguments) < 1 {
		event.Respond("Usage: welcome <channel> [message]")
		return
	}

	targetChan := services.GetChannel(event.Arguments[0])
	if targetChan == nil || targetChan.Name() == "" {
		event.Respond("Channel is invalid")
		return
	}

	if !targetChan.HasOp(event.Connection.GetNick()) {
		event.Respond("I do not have permissions in that channel")
	}

	if len(event.Arguments) == 1 {
		mapping[targetChan.Name()] = ""
		event.Respond("Channel's welcome message has been cleared")
	} else {
		mapping[targetChan.Name()] = strings.Join(event.Arguments[1:], " ")
		event.Respond("Channel's welcome message has been set to: " + mapping[targetChan.Name()])
	}
}

func runJoin(event events.Join) {
	message := mapping[event.Channel.Name()]

	if message == "" {
		return
	}

	event.Connection.Notice(event.User.Nickname(), message)
}
