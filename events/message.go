package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Message struct {
	Message    string
	User       core.User
	Channel    core.Channel
	Connection *irc.Connection
}

func (m *Message) EventName() string {
	return "message"
}
