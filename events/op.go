package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Op struct {
	Connection *irc.Connection
	Channel    core.Channel
	Target     core.User
	Sender     core.User
}

func (o *Op) EventName() string {
	return "op"
}
