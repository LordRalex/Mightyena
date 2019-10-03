package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Devoice struct {
	Event

	Connection *irc.Connection
	Channel    core.Channel
	Target     core.User
	Sender     core.User
}

func (d *Devoice) EventName() string {
	return "devoice"
}
