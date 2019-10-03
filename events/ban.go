package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Ban struct {
	Event

	Connection *irc.Connection
	Mask       string
	Banner     core.User
	Channel    core.Channel
}

func (b *Ban) EventName() string {
	return "ban"
}
