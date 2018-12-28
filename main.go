package main

import (
	"github.com/lordralex/mightyena/config"
	"github.com/lordralex/mightyena/core"
	"github.com/lordralex/mightyena/database"
	"github.com/thoj/go-ircevent"
)

var bot *irc.Connection

func main() {
	coreConfig, err := config.Get("core", "json")
	if err != nil {
		panic(err)
		return
	}

	dbUrl, err := coreConfig.GetString("database")
	if err != nil {
		panic(err)
		return
	}

	err = database.CreatePool(dbUrl)
	if err != nil {
		panic(err)
		return
	}

	nick, _ := coreConfig.GetString("nickname")
	server, _ := coreConfig.GetString("server")

	bot = irc.IRC(nick, "Bot")
	bot.Debug = true
	bot.UseTLS = true
	bot.RequestCaps = append(bot.RequestCaps, "account-notify")

	core.CreateServiceHandlers(bot)
	err = bot.Connect(server)
	if err != nil {
		panic(err)
		return
	}
	bot.Loop()
}

func Shutdown() {
	bot.Disconnect()
}