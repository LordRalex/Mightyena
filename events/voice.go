package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Voice struct {
	Connection *irc.Connection
	Channel    core.Channel
	Target     core.User
	Sender     core.User
}

func (v *Voice) EventName() string {
	return "voice"
}
