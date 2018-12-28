package event

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Action struct {
	action     string
	user       core.User
	channel    core.Channel
	connection *irc.Connection
}

func (a *Action) GetConnection() *irc.Connection {
	return a.connection
}

func (a *Action) GetAction() string {
	return a.action
}

func (a *Action) GetUser() core.User {
	return a.user
}

func (a *Action) GetChannel() core.Channel {
	return a.channel
}
