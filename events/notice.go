package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Notice struct {
	Event

	Message    string
	User       core.User
	Channel    core.Channel
	Connection *irc.Connection
}

func (n *Notice) EventName() string {
	return "notice"
}
