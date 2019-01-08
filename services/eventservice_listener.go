package services

import (
	"github.com/lordralex/mightyena/events"
	"github.com/thoj/go-ircevent"
	"strings"
)

func fireMessageEvent(event *irc.Event) {
	message := event.Message()
	nick := event.Nick
	channel := event.Arguments[0]

	if !strings.HasPrefix(channel, "#") {
		channel = ""
	}

	evt := events.CreateMessageEvent(event.Connection, message, GetUser(nick), GetChannel(channel))
	fireEvent(evt)
}
