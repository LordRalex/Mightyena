package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Ban interface {
	Event

	Channel() core.Channel
	Banner() core.User
	Mask() string
}

type ban struct {
	connection *irc.Connection

	mask string
	banner core.User
	channel core.Channel
}

func CreateBanEvent(connection *irc.Connection, mask string, banner core.User, channel core.Channel) Ban {
	return &ban{connection: connection, mask: mask, banner: banner, channel: channel}
}

func (b *ban) Channel() core.Channel {
	return b.channel
}

func (b *ban) Mask() string {
	return b.mask
}

func (b *ban) Banner() core.User {
	return b.banner
}

func (b *ban) Connection() *irc.Connection {
	return b.connection
}

func (b *ban) EventName() string {
	return "ban"
}