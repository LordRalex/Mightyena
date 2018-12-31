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

func handleNamesEventChannelService(event *irc.Event) {
	//discard first 2 args
	//third arg is the channel
	//rest are users
	channelName := event.Arguments[2]

	if namesBuffer[channelName] == nil {
		namesBuffer[channelName] = make([]string, 0)
	}

	namesBuffer[channelName] = append(namesBuffer[channelName], event.Arguments[3:]...)
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
		if strings.HasPrefix(v, "@") {
			v = strings.TrimPrefix(v, "@")
			channel.ops = append(channel.ops, v)
		} else if strings.HasPrefix(v, "+") {
			v = strings.TrimPrefix(v, "+")
			channel.voiced = append(channel.voiced, v)
		}
		channel.users = append(channel.users, v)
		event.Connection.Whois(v)
	}

	channelWriter.Lock()
	defer channelWriter.Unlock()
	channelCache[channelName] = channel
	namesBuffer[channelName] = nil

	logging.GetLogger("CHAN SERVICE").Log(logging.Debug, "%+v", channelCache)
}
