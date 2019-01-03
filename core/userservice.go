package core

import (
	"github.com/lordralex/mightyena/logging"
	"github.com/thoj/go-ircevent"
	"sync"
	"time"
)

var userCache = make(map[string]*user)

var userWriter = sync.RWMutex{}

var userServiceLogger = logging.GetLogger("USER SERVICE")

func GetUser(nickname string) (User, error) {
	return getUser(nickname), nil
}

func getUser(nickname string) *user {
	userWriter.RLock()
	defer userWriter.RUnlock()
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

	userWriter.Lock()
	defer userWriter.Unlock()
	userCache[user.GetNickname()] = user
}

func handleQuitEventUserService(event *irc.Event) {
	//user never was in the cache, should not happen though...
	if u := getUser(event.Nick); u == nil {
		userServiceLogger.Log(logging.Info, "User not in the cache when they quit: [%s]", event.Nick)
		return
	}

	userWriter.Lock()
	defer userWriter.Unlock()
	userCache[event.Nick] = nil

	if event.Nick == event.Connection.GetNick() {
		userCache = make(map[string]*user)
	}
}

func handleNickEventUserService(event *irc.Event) {
	//user never was in the cache, should not happen though...
	var user *user
	if user = getUser(event.Nick); user == nil {
		userServiceLogger.Log(logging.Info,"User not in the cache when they changed nicks: [%s]", event.Nick)
		return
	}

	userWriter.Lock()
	defer userWriter.Unlock()

	//update cache with new nickname
	user.nickname = event.Arguments[0]
	userCache[event.Nick] = nil
	userCache[user.GetNickname()] = user
}

func handlePartEventUserService(event *irc.Event) {
	//determine if they don't exist in any channels we track
	for _, v := range channelCache {
		for _, u := range v.users {
			if u == event.Nick {
				return
			}
		}
	}

	userWriter.Lock()
	defer userWriter.Unlock()

	userCache[event.Nick] = nil
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
	userWriter.Lock()
	defer userWriter.Unlock()

	userServiceLogger.Log(logging.Debug, "Running user service cleanup")

	newListing := make(map[string]*user)

	for k, v := range userCache {
		if v != nil {
			newListing[k] = v
		}
	}

	userCache = newListing
}