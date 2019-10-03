package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Command struct {
	Event

	Connection *irc.Connection
	Command    string
	Arguments  []string
	User       core.User
	Channel    core.Channel
}

func (c *Command) EventName() string {
	return "command"
}

func (c *Command) Respond(message string) {
	if c.Channel != nil {
		c.Connection.Privmsg(c.Channel.Name(), message)
	} else {
		c.Connection.Privmsg(c.User.Nickname(), message)
	}
}
