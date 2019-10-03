package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Part struct {
	Connection *irc.Connection
	User       core.User
	Channel    core.Channel
	Message    string
}

func (p *Part) EventName() string {
	return "part"
}
