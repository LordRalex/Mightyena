package services

import (
	"github.com/lordralex/mightyena/core"
	"github.com/lordralex/mightyena/logging"
	"github.com/thoj/go-ircevent"
	"sync"
	"time"
)

var userCache = make(map[string]*user)

var userWriter = sync.RWMutex{}

var userLogger = logging.GetLogger("USER-SVC")

func GetUser(nickname string) core.User {
	return getUser(nickname)
}

func getUser(nickname string) *user {
	userWriter.RLock()
	defer userWriter.RUnlock()
	return userCache[nickname]
}

func handleJoinEventUserService(event *irc.Event) {
	//user already exists in the cache, don't need to process
	if u := getUser(event.Nick); u != nil {
		return
	}

	user := &user{
		nickname:     event.Nick,
		host:         event.Host,
		loginName:    event.User,
		nickservName: "",
	}

	userWriter.Lock()
	defer userWriter.Unlock()
	userCache[user.Nickname()] = user
}

func handleQuitEventUserService(event *irc.Event) {
	//user never was in the cache, should not happen though...
	if u := getUser(event.Nick); u == nil {
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
		return
	}

	userWriter.Lock()
	defer userWriter.Unlock()

	//update cache with new nickname
	user.nickname = event.Arguments[0]
	userCache[event.Nick] = nil
	userCache[user.Nickname()] = user
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

func handleWhoEventUserService(event *irc.Event) {
	nick := event.Arguments[5]
	userName := event.Arguments[2]
	host := event.Arguments[3]

	u := userCache[nick]

	if u == nil {
		u = &user{}
	}

	u.nickname = nick
	u.loginName = userName
	u.host = host

	userWriter.Lock()
	defer userWriter.Unlock()
	userCache[nick] = u

	userLogger.Debug("Added to cache: %+v", userCache[nick])
}

func startCleanupUserService() {
	ticker := time.NewTicker(5 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				userServiceTick()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func userServiceTick() {
	userWriter.Lock()
	defer userWriter.Unlock()

	newListing := make(map[string]*user)

	for k, v := range userCache {
		if v != nil {
			newListing[k] = v
		}
	}

	userCache = newListing
}
