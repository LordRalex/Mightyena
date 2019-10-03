package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Join struct {
	Event

	Connection *irc.Connection
	User       core.User
	Channel    core.Channel
}

func (j *Join) EventName() string {
	return "join"
}
