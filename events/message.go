package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Message interface {
	Event

	Message() string
	User() core.User
	Channel() core.Channel
}

type message struct {
	message    string
	user       core.User
	channel    core.Channel
	connection *irc.Connection
}

func CreateMessageEvent(connection *irc.Connection, msg string, user core.User, channel core.Channel) Message {
	return &message{connection: connection, user: user, message: msg, channel: channel}
}

func (m *message) EventName() string {
	return "message"
}

func (m *message) Connection() *irc.Connection {
	return m.connection
}

func (m *message) User() core.User {
	return m.user
}

func (m *message) Channel() core.Channel {
	return m.channel
}

func (m *message) Message() string {
	return m.message
}
