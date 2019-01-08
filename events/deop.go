package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Deop interface {
	Event

	Target() core.User
	Sender() core.User
	Channel() core.Channel
}

type deop struct {
	connection *irc.Connection

	channel core.Channel
	target core.User
	sender core.User
}

func CreateDeopEvent(connection *irc.Connection, channel core.Channel, target core.User, sender core.User) Deop {
	return &deop{connection: connection, channel: channel, target: target, sender: sender}
}

func (d *deop) Target() core.User {
	return d.target
}

func (d *deop) Sender() core.User {
	return d.sender
}

func (d *deop) Channel() core.Channel {
	return d.channel
}

func (d *deop) Connection() *irc.Connection {
	return d.connection
}

func (d *deop) EventName() string {
	return "deop"
}