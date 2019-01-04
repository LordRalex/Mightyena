package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Action interface {
	Event

	Action() string
	User() core.User
	Channel() core.Channel
}

type action struct {
	action     string
	user       core.User
	channel    core.Channel
	connection *irc.Connection
}

func (a *action) Connection() *irc.Connection {
	return a.connection
}

func (a *action) Action() string {
	return a.action
}

func (a *action) User() core.User {
	return a.user
}

func (a *action) Channel() core.Channel {
	return a.channel
}

func (a *action) EventName() string {
	return "action"
}