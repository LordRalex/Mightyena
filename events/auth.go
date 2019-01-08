package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Auth interface {
	Event

	Account() string
	User() core.User
}

type auth struct {
	connection *irc.Connection
	account    string
	user       core.User
}

func CreateAuthEvent(connection *irc.Connection, account string, user core.User) Auth {
	return &auth{connection: connection, account: account, user: user}
}

func (a *auth) Account() string {
	return a.account
}

func (a *auth) Connection() *irc.Connection {
	return a.connection
}

func (a *auth) User() core.User {
	return a.user
}

func (a *auth) EventName() string {
	return "auth"
}
