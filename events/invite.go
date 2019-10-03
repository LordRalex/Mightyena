package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Invite struct {
	Connection *irc.Connection
	Channel    string
	User       core.User
}

func (i *Invite) EventName() string {
	return "invite"
}
