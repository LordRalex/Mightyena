package events

import "github.com/thoj/go-ircevent"

type Event interface {
	EventName() string

	Connection() *irc.Connection
}
