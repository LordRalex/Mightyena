package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Auth struct {
	Connection *irc.Connection
	Account    string
	User       core.User
}

func (a *Auth) EventName() string {
	return "auth"
}
