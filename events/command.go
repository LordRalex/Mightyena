package events

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
)

type Command interface {
	Event

	Command() string
	Arguments() []string
	User() core.User
	Channel() core.Channel

	Respond(message string)
}

type command struct {
	connection *irc.Connection
	command    string
	arguments  []string
	user       core.User
	channel    core.Channel
}

func CreateCommandEvent(connection *irc.Connection, cmd string, arguments []string, user core.User, channel core.Channel) Command {
	return &command{connection: connection, command: cmd, arguments: arguments, user: user, channel: channel}
}

func (c *command) Channel() core.Channel {
	return c.channel
}

func (c *command) User() core.User {
	return c.user
}

func (c *command) Arguments() []string {
	return c.arguments
}

func (c *command) Command() string {
	return c.command
}

func (c *command) Connection() *irc.Connection {
	return c.connection
}

func (c *command) EventName() string {
	return "command"
}

func (c *command) Respond(message string) {
	if c.Channel() != nil {
		c.Connection().Privmsg(c.Channel().Name(), message)
	} else {
		c.Connection().Privmsg(c.User().Nickname(), message)
	}
}
