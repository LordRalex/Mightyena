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

	evt := events.CreateMessageEvent(event.Connection, message, GetUser(nick), GetChannel(channel))
	fireEvent(evt)

	//since messages can be events, we'll bounce off this one

	if strings.HasPrefix(message, commandPrefix) {
		fireCommandEvent(evt)
	}
}

func fireCommandEvent(event events.Message) {
	parts := strings.Split(event.Message(), " ")
	parts[0] = strings.TrimPrefix(parts[0], commandPrefix)

	evt := events.CreateCommandEvent(event.Connection(), parts[0], parts[1:], event.User(), event.Channel())
	fireEvent(evt)

	executeCommand(evt)
}
