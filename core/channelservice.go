package core

import (
	"github.com/lordralex/mightyena/logging"
	"github.com/thoj/go-ircevent"
	"strings"
	"sync"
)

var channelCache = make(map[string]*channel)
var namesBuffer = make(map[string][]string)

var channelWriter = sync.RWMutex{}

var chanServiceLogger = logging.GetLogger("CHAN SERVICE")

func handleNamesEventChannelService(event *irc.Event) {
	//discard first 2 args
	//third arg is the channel
	//rest are users
	channelName := event.Arguments[2]
	names := strings.Split(event.Arguments[3], " ")

	if namesBuffer[channelName] == nil {
		namesBuffer[channelName] = make([]string, 0)
	}

	namesBuffer[channelName] = append(namesBuffer[channelName], names...)
}

func handleNamesEndEventChannelService(event *irc.Event) {
	channelName := event.Arguments[1]
	names := namesBuffer[channelName]

	channel := &channel{
		users:  make([]string, 0),
		name:   channelName,
		ops:    make([]string, 0),
		voiced: make([]string, 0),
	}

	for _, v := range names {
		name := v

		if strings.HasPrefix(name, ":") {
			name = strings.TrimPrefix(name, ":")
		}

		if strings.HasPrefix(name, "@") {
			name = strings.TrimPrefix(name, "@")
			channel.ops = append(channel.ops, name)
		} else if strings.HasPrefix(name, "+") {
			name = strings.TrimPrefix(name, "+")
			channel.voiced = append(channel.voiced, name)
		}
		channel.users = append(channel.users, name)
		event.Connection.Whois(name)
	}

	channelWriter.Lock()
	defer channelWriter.Unlock()
	channelCache[channelName] = channel
	namesBuffer[channelName] = nil

	chanServiceLogger.Log(logging.Debug, "Channel: %s", channelName)
	for _, v := range namesBuffer[channelName] {
		chanServiceLogger.Log(logging.Debug, "  User: %s", v)
	}
}

func handlePartEventChannelService(event *irc.Event) {
	nickname := event.Nick
	channelName := event.Arguments[0]

	if event.Nick == event.Connection.GetNick() {
		channelWriter.Lock()
		defer channelWriter.Unlock()

		channelCache[channelName] = nil
		newList := make(map[string]*channel)
		for k, v := range newList {
			if v != nil {
				newList[k] = v
			}
		}
		channelCache = newList
	} else {
		ch := channelCache[channelName]
		ch.removeUser(nickname)
	}
}

func handleQuitEventChannelService(event *irc.Event) {
	nick := event.Nick

	//remove user from all channels we know of
	for _, v := range channelCache {
		v.removeUser(nick)
	}

	//reset cache if we were the one that quit
	if event.Nick == event.Connection.GetNick() {
		channelCache = make(map[string]*channel)
	}
}