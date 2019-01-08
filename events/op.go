package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Op interface {
	Event

	Target() core.User
	Sender() core.User
	Channel() core.Channel
}

type op struct {
	connection *irc.Connection

	channel core.Channel
	target core.User
	sender core.User
}

func CreateOpEvent(connection *irc.Connection, channel core.Channel, target core.User, sender core.User) Op {
	return &op{connection: connection, channel: channel, target: target, sender: sender}
}

func (o *op) Target() core.User {
	return o.target
}

func (o *op) Sender() core.User {
	return o.sender
}

func (o *op) Channel() core.Channel {
	return o.channel
}

func (o *op) Connection() *irc.Connection {
	return o.connection
}

func (o *op) EventName() string {
	return "op"
}