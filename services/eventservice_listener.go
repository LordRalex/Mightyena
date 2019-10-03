package services

import (
	"github.com/lordralex/mightyena/config"
	"github.com/lordralex/mightyena/events"
	"github.com/thoj/go-ircevent"
	"strings"
)

var commandPrefix = "!!"

func init() {
	cfg, err := config.Get("core", "env")
	if err == nil && cfg != nil {
		prefix, err := cfg.GetString("commandPrefix")
		if err == nil {
			prefix = strings.TrimSpace(prefix)
			if prefix != "" {
				commandPrefix = prefix
			}
		}
	}
}

func fireMessageEvent(event *irc.Event) {
	message := event.Message()
	nick := event.Nick
	channel := event.Arguments[0]

	if !strings.HasPrefix(channel, "#") {
		channel = ""
	}

	evt := &events.Message{
		Message:    message,
		User:       GetUser(nick),
		Channel:    GetChannel(channel),
		Connection: event.Connection,
	}
	fireEvent(evt)

	//since messages can be events, we'll bounce off this one

	if strings.HasPrefix(message, commandPrefix) {
		fireCommandEvent(evt)
	}
}

func fireCommandEvent(event *events.Message) {
	parts := strings.Split(event.Message, " ")
	if len(parts) < 2 {
		return
	}
	parts[0] = strings.TrimPrefix(parts[0], commandPrefix)

	evt := &events.Command{
		Connection: event.Connection,
		Command:    parts[0],
		Arguments:  parts[1:],
		User:       event.User,
		Channel:    event.Channel,
	}
	fireEvent(evt)

	executeCommand(evt)
}
