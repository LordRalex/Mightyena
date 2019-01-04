package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

func CreateMessageEvent(msg string, user core.User, channel core.Channel, connection *irc.Connection) Message {
	return &message{
		connection: connection,
		user:       user,
		message:    msg,
		channel:    channel,
	}
}
