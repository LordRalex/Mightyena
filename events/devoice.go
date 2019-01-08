package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Devoice interface {
	Event

	Target() core.User
	Sender() core.User
	Channel() core.Channel
}

type devoice struct {
	connection *irc.Connection

	channel core.Channel
	target core.User
	sender core.User
}

func CreateDevoiceEvent(connection *irc.Connection, channel core.Channel, target core.User, sender core.User) Devoice {
	return &devoice{connection: connection, channel: channel, target: target, sender: sender}
}

func (d *devoice) Target() core.User {
	return d.target
}

func (d *devoice) Sender() core.User {
	return d.sender
}

func (d *devoice) Channel() core.Channel {
	return d.channel
}

func (d *devoice) Connection() *irc.Connection {
	return d.connection
}

func (d *devoice) EventName() string {
	return "devoice"
}