package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Quit struct {
	Event

	Connection *irc.Connection
	User       core.User
	Message    string
}

func (q *Quit) EventName() string {
	return "quit"
}
