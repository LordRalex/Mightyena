package main

import (
	"github.com/lordralex/mightyena/config"
	"github.com/lordralex/mightyena/database"
	"github.com/lordralex/mightyena/listeners"
	"github.com/lordralex/mightyena/logging"
	"github.com/lordralex/mightyena/services"
	"github.com/thoj/go-ircevent"
	"os"
	"time"
)

var bot *irc.Connection

var logger = logging.GetLogger("CORE")

func main() {
	coreConfig, err := config.Get("core", "env")
	if err != nil {
		logger.Error(err.Error())
		return
	}

	dbUrl, err := coreConfig.GetString("database")
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = database.CreatePool(dbUrl)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	nick, _ := coreConfig.GetString("nickname")
	server, _ := coreConfig.GetString("server")

	bot = irc.IRC(nick, "Bot")
	bot.Debug = true
	bot.UseTLS = true
	bot.RequestCaps = append(bot.RequestCaps, "account-notify")

	services.CreateServiceHandlers(bot)
	err = bot.Connect(server)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	if len(os.Args) > 1 {
		go func() {
			time.Sleep(3 * time.Second)
			for _, v := range os.Args[1:] {
				logger.Info("Joining %s", v)
				bot.Join(v)
			}
		}()
	} else {
		logger.Warning("No channel specified")
	}

	//test message event listener
	listeners.RegisterTest()
	listeners.RegisterModules()

	bot.Loop()
}