package webclient

import (
	"github.com/lordralex/mightyena/config"
	"github.com/lordralex/mightyena/events"
	"github.com/lordralex/mightyena/logging"
	"github.com/lordralex/mightyena/services"
	"regexp"
)

const ModuleName = "webclient"

var Logger = logging.GetLogger(ModuleName)

func Load() {
	services.RegisterJoin(ModuleName, runJoin)
}

func runJoin(event *events.Join) {
	cfg, err := config.Get(ModuleName, "mysql")
	if err != nil {
		Logger.Error("Error loading config: %s", err.Error())
		return
	}

	channels, err := cfg.GetStringList("channels")
	if err != nil {
		Logger.Error("Error getting channel list: %s", err.Error())
		return
	}

	protected := false
	for _, c := range channels {
		if c == event.Channel.Name() {
			protected = true
			break
		}
	}

	if !protected {
		return
	}

	masks, err := cfg.GetStringList("masks")
	if err != nil {
		Logger.Error("Error getting channel list: %s", err.Error())
		return
	}

	conn := event.User.Nickname() + "!" + event.User.LoginName() + "@" + event.User.Hostname()
	for _, m := range masks {
		reg, err := regexp.Compile(m)
		if err != nil {
			Logger.Error("Failed compiling regex (%s): %s", m, err.Error())
			continue
		}
		if reg.MatchString(conn) {
			msg, err := cfg.GetString("kick")
			if err != nil || msg == "" {
				msg = "Client not permitted"
			}
			banMask := "*!" + event.User.LoginName() + "@" + event.User.Hostname()
			event.Connection.Mode(event.Channel.Name(), "+b "+banMask)
			event.Connection.Kick(event.User.Nickname(), event.Channel.Name(), msg)
			break
		}
	}
}
