package main

import (
	"github.com/lordralex/mightyena/config"
	"github.com/lordralex/mightyena/core"
	"github.com/lordralex/mightyena/database"
	"github.com/lordralex/mightyena/logging"
	"github.com/thoj/go-ircevent"
	"os"
	"time"
)

var bot *irc.Connection

var logger = logging.GetLogger("CORE")

func main() {
	coreConfig, err := config.Get("core", "json")
	if err != nil {
		logger.Log(logging.Error, err.Error())
		return
	}

	dbUrl, err := coreConfig.GetString("database")
	if err != nil {
		logger.Log(logging.Error, err.Error())
		return
	}

	err = database.CreatePool(dbUrl)
	if err != nil {
		logger.Log(logging.Error, err.Error())
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
		logger.Log(logging.Error, err.Error())
		return
	}

	if len(os.Args) > 1 {
		go func() {
			time.Sleep(5 * time.Second)
			bot.Join(os.Args[1])
		}()
	}

	bot.Loop()
}

func Shutdown() {
	bot.Disconnect()
}