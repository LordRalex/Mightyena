package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Unban interface {
	Event

	Channel() core.Channel
	Banner() core.User
	Mask() string
}

type unban struct {
	connection *irc.Connection

	mask string
	banner core.User
	channel core.Channel
}

func CreateUnbanEvent(connection *irc.Connection, mask string, banner core.User, channel core.Channel) Unban {
	return &unban{connection: connection, mask: mask, banner: banner, channel: channel}
}

func (u *unban) Channel() core.Channel {
	return u.channel
}

func (u *unban) Mask() string {
	return u.mask
}

func (u *unban) Banner() core.User {
	return u.banner
}

func (u *unban) Connection() *irc.Connection {
	return u.connection
}

func (u *unban) EventName() string {
	return "unban"
}