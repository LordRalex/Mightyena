package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Invite interface {
	Event

	Channel() string
	User() core.User
}

type invite struct {
	connection *irc.Connection
	channel string
	user core.User
}

func CreateInviteEvent(connection *irc.Connection, channel string, user core.User) Invite {
	return &invite{connection: connection, channel: channel, user: user}
}

func (i *invite) User() core.User {
	return i.user
}

func (i *invite) Channel() string {
	return i.channel
}

func (i *invite) Connection() *irc.Connection {
	return i.connection
}

func (i *invite) EventName() string {
	return "invite"
}