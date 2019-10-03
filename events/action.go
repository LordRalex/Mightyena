package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Action struct {
	Action     string
	User       core.User
	Channel    core.Channel
	Connection *irc.Connection
}

func (a *Action) EventName() string {
	return "action"
}
