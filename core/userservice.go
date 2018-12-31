package core

import (
	"github.com/lordralex/mightyena/logging"
	"github.com/thoj/go-ircevent"
	"sync"
	"time"
)

var userCache = make(map[string]*user)

var writeInProgress = sync.RWMutex{}

func GetUser(nickname string) (User, error) {
	return getUser(nickname), nil
}

func getUser(nickname string) *user {
	writeInProgress.RLock()
	defer writeInProgress.RUnlock()
	return userCache[nickname]
}

func handleJoinEventUserService(event *irc.Event) {
	//user already exists in the cache, don't need to process
	if u  := getUser(event.Nick); u != nil {
		return
	}

	user := &user{
		nickname: event.Nick,
		host: event.Host,
		loginName: event.User,
		nickservName: "",
	}

	writeInProgress.Lock()
	defer writeInProgress.Unlock()
	userCache[user.GetNickname()] = user
}

func handleQuitEventUserService(event *irc.Event) {
	//user never was in the cache, should not happen though...
	if u := getUser(event.Nick); u == nil {
		logging.GetLogger("USER SERVICE").Log(logging.Info, "User not in the cache when they quit: [%s]", event.Nick)
		return
	}

	writeInProgress.Lock()
	defer writeInProgress.Unlock()
	userCache[event.Nick] = nil
}

func handleNickEventUserService(event *irc.Event) {
	//user never was in the cache, should not happen though...
	var user *user
	if user = getUser(event.Nick); user == nil {
		logging.GetLogger("USER SERVICE").Log(logging.Info,"User not in the cache when they changed nicks: [%s]", event.Nick)
		return
	}

	writeInProgress.Lock()
	defer writeInProgress.Unlock()

	//update cache with new nickname
	user.nickname = event.Arguments[0]
	userCache[event.Nick] = nil
	userCache[user.GetNickname()] = user
}

func startCleanupUserService() {
	ticker := time.NewTicker(5 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <- ticker.C:
				userServiceTick()
			case <- quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func userServiceTick() {
	writeInProgress.Lock()
	defer writeInProgress.Unlock()

	newListing := make(map[string]*user)

	for k, v := range userCache {
		if v != nil {
			newListing[k] = v
		}
	}

	userCache = newListing
}