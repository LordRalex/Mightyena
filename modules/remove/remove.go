package remove

import (
	"github.com/lordralex/mightyena/config"
	"github.com/lordralex/mightyena/events"
	"github.com/lordralex/mightyena/logging"
	"github.com/lordralex/mightyena/services"
	"strings"
)

const ModuleName = "remove"

var Logger = logging.GetLogger(ModuleName)

func Load() {
	services.RegisterCommand(ModuleName, "kick", runKick)
	services.RegisterCommand(ModuleName, "ban", runBan)

	services.RegisterCommand(ModuleName, "remove", runKick)
	services.RegisterCommand(ModuleName, "terminate", runBan)
}

func runKick(event events.Command) {
	handle(event, false)
}

func runBan(event events.Command) {
	handle(event, true)
}

func handle(event events.Command, ban bool) {
	if event.Channel() == nil {
		return
	}

	if len(event.Arguments()) == 0 {
		event.Respond("Username required")
		return
	}

	if !event.Channel().HasVoiceOrOp(event.User().Nickname()) {
		return
	}

	cfg, err := config.Get(ModuleName, "mysql")
	if err != nil {
		Logger.Error("Error getting config: %s", err.Error())
		return
	}

	channels, err := cfg.GetStringList("channels")
	if err != nil {
		Logger.Error("Error getting config: %s", err.Error())
		return
	}

	valid := false

	for _, c := range channels {
		if c == event.Channel().Name() {
			valid = true
			break
		}
	}

	if !valid {
		return
	}

	msg := strings.Join(event.Arguments()[1:], " ")
	if msg == "" {
		msg = "You have been removed from this channel"
	}

	event.Connection().Kick(event.User().Nickname(), event.Channel().Name(), msg)
	if ban {
		mask := "*!*@" + event.User().Hostname()
		event.Connection().Mode(event.Channel().Name(), "+b "+mask)
	}
}
