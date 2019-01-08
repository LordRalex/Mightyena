package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Part interface {
	Event

	User() core.User
	Channel() core.Channel
	Message() string
}

type part struct {
	connection *irc.Connection
	user core.User
	channel core.Channel
	message string
}

func (p *part) Message() string {
	return p.message
}

func (p *part) Channel() core.Channel {
	return p.channel
}

func (p *part) User() core.User {
	return p.user
}

func (p *part) Connection() *irc.Connection {
	return p.connection
}

func (p *part) EventName() string {
	return "part"
}
