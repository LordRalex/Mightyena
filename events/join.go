package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Join interface {
	Event

	User() core.User
	Channel() core.Channel
}

type join struct {
	connection *irc.Connection
	user core.User
	channel core.Channel
}

func CreateJoinEvent(connection *irc.Connection, user core.User, channel core.Channel) Join {
	return &join{connection: connection, user: user, channel: channel}
}

func (j *join) Channel() core.Channel {
	return j.channel
}

func (j *join) Connection() *irc.Connection {
	return j.connection
}

func (j *join) User() core.User {
	return j.user
}

func (j *join) EventName() string {
	return "join"
}