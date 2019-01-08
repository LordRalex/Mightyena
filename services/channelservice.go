package services

import (
	"github.com/lordralex/mightyena/core"
	"github.com/thoj/go-ircevent"
	"strings"
	"sync"
)

var channelCache = make(map[string]*channel)
var namesBuffer = make(map[string][]string)

var channelWriter = sync.RWMutex{}

func GetChannel(chanName string) core.Channel {
	return getChannel(chanName)
}

func getChannel(chanName string) *channel {
	if chanName == "" {
		return nil
	}
	channelWriter.RLock()
	defer channelWriter.RUnlock()
	return channelCache[chanName]
}

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
	}

	channelWriter.Lock()
	defer channelWriter.Unlock()
	channelCache[channelName] = channel

	event.Connection.SendRaw("WHO " + channelName)
}

func handlePartEventChannelService(event *irc.Event) {
	nickname := event.Nick
	channelName := event.Arguments[0]

	channelWriter.Lock()
	defer channelWriter.Unlock()
	if event.Nick == event.Connection.GetNick() {
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

	channelWriter.Lock()
	defer channelWriter.Unlock()

	//remove user from all channels we know of
	for _, v := range channelCache {
		v.removeUser(nick)
	}

	//reset cache if we were the one that quit
	if event.Nick == event.Connection.GetNick() {
		channelCache = make(map[string]*channel)
	}
}
