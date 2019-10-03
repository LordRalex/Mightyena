package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Kick struct {
	Event

	Connection *irc.Connection
	Sender     core.User
	Target     core.User
	Channel    core.Channel
	Message    string
}

func (k *Kick) EventName() string {
	return "kick"
}
