package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Voice interface {
	Event

	Target() core.User
	Sender() core.User
	Channel() core.Channel
}

type voice struct {
	connection *irc.Connection

	channel core.Channel
	target core.User
	sender core.User
}

func CreateVoiceEvent(connection *irc.Connection, channel core.Channel, target core.User, sender core.User) Voice {
	return &voice{connection: connection, channel: channel, target: target, sender: sender}
}

func (v *voice) Target() core.User {
	return v.target
}

func (v *voice) Sender() core.User {
	return v.sender
}

func (v *voice) Channel() core.Channel {
	return v.channel
}

func (v *voice) Connection() *irc.Connection {
	return v.connection
}

func (v *voice) EventName() string {
	return "voice"
}