package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Kick interface {
	Event

	Sender() core.User
	Target() core.User
	Message() string
	Channel() core.Channel
}

type kick struct {
	connection *irc.Connection
	sender core.User
	target core.User
	channel core.Channel
	message string
}

func CreateKickEvent(connection *irc.Connection, sender core.User, target core.User, channel core.Channel, message string) Kick {
	return &kick{connection: connection, sender: sender, target: target, channel: channel, message: message}
}

func (k *kick) Message() string {
	return k.message
}

func (k *kick) Channel() core.Channel {
	return k.channel
}

func (k *kick) Target() core.User {
	return k.target
}

func (k *kick) Sender() core.User {
	return k.sender
}

func (k *kick) Connection() *irc.Connection {
	return k.connection
}

func (k *kick) EventName() string {
	return "kick"
}