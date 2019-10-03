package antispam

import (
	"github.com/lordralex/mightyena/events"
	"sync"
	"time"
)

var messageHistory = make(map[string][]*messageData, 0)
var locker = sync.Mutex{}

type messageData struct {
	message string
	time    time.Time
}

func Load() {
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		for {
			select {
			case <-ticker.C:
				updateCache()
			}
		}
	}()
}

func handleMessageEvent(event events.Message) {
	locker.Lock()
	defer locker.Unlock()

	data, exists := messageHistory[event.User.Nickname()]
	if !exists {
		data = make([]*messageData, 0)
	}

	data = append(data, &messageData{message: event.Message, time: time.Now()})
	messageHistory[event.User.Nickname()] = data
}

func handleKickEvent(kick events.Kick) {
	deleteFromCache(kick.Target.Nickname())
}

func handlePartEvent(part events.Part) {
	deleteFromCache(part.User.Nickname())
}

func handleQuitEvent(part events.Part) {
	deleteFromCache(part.User.Nickname())
}

func deleteFromCache(key string) {
	locker.Lock()
	defer locker.Unlock()

	delete(messageHistory, key)
}

func updateCache() {
	locker.Lock()
	defer locker.Unlock()

	for k, v := range messageHistory {
		for k2, v2 := range v {
			//removes any messages that are over an hour old
			if v2.time.Add(time.Hour).Before(time.Now()) {
				v = append(v[:k2], v[k2+1:]...)
			}
		}

		//if the resulting cache is empty, let's just delete it from the main cache
		if len(v) == 0 {
			delete(messageHistory, k)
		}
	}
}
