package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Notice interface {
	Event

	Message() string
	User() core.User
	Channel() core.Channel
}

type notice struct {
	message    string
	user       core.User
	channel    core.Channel
	connection *irc.Connection
}

func CreateNoticeEvent(connection *irc.Connection, msg string, user core.User, channel core.Channel) Notice {
	return &notice{connection: connection, user: user, message: msg, channel: channel}
}

func (n *notice) EventName() string {
	return "notice"
}

func (n *notice) Connection() *irc.Connection {
	return n.connection
}

func (n *notice) User() core.User {
	return n.user
}

func (n *notice) Channel() core.Channel {
	return n.channel
}

func (n *notice) Message() string {
	return n.message
}