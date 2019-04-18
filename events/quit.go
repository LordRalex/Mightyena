package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Quit interface {
	Event

	User() core.User
	Channel() core.Channel
	Message() string
}

type quit struct {
	connection *irc.Connection
	user core.User
	message string
}

func (q *quit) Message() string {
	return q.message
}

func (q *quit) User() core.User {
	return q.user
}

func (q *quit) Connection() *irc.Connection {
	return q.connection
}

func (q *quit) EventName() string {
	return "quit"
}
