package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Unban struct {
	Connection *irc.Connection
	Mask       string
	Banner     core.User
	Channel    core.Channel
}

func (u *Unban) EventName() string {
	return "unban"
}
