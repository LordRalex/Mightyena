package factoid

import (
	"github.com/lordralex/mightyena/core"
	"github.com/lordralex/mightyena/database"
	"github.com/lordralex/mightyena/events"
	"github.com/lordralex/mightyena/format"
	"github.com/lordralex/mightyena/services"
	"strings"
)

const ModuleName = "factoids"
const messageFormat = format.IrcBold + "%s:" + format.IrcPlain + " %s"

func Load() {
	services.RegisterCommand(ModuleName, ">", handleToUser)
	services.RegisterCommand(ModuleName, "<", handleToSelf)
	services.RegisterCommand(ModuleName, ".", handleToChannel)
	services.RegisterCommand(ModuleName, "", handleToChannel)
}

func handleToUser(event events.Command) {
	if event.Channel() == nil {
		return
	}

	handle(event, event.Arguments()[0], event.Arguments()[1], event.Channel(), event.User())
}

func handleToSelf(event events.Command) {
	handle(event, event.Arguments()[0], event.Arguments()[0], nil, event.User())
}

func handleToChannel(event events.Command) {
	handle(event, event.Arguments()[0], event.Arguments()[0], event.Channel(), event.User())
}

func handle(event events.Command, prefix, key string, channel core.Channel, user core.User) {
	factoidInfo := getFactoid(key)

	if factoidInfo == nil || len(factoidInfo) == 0 {
		event.Connection().Noticef(user.Nickname(), "No factoid with name (%s) found", key)
		return
	}

	var target string
	if channel != nil {
		target = channel.Name()
		for _, v := range factoidInfo {
			event.Connection().Privmsgf(target, messageFormat, prefix, format.ParseFromBBCode(v))
		}
	} else {
		target = user.Nickname()
		for _, v := range factoidInfo {
			event.Connection().Noticef(target, messageFormat, prefix, format.ParseFromBBCode(v))
		}
	}
}

func getFactoid(key string) []string {
	db := database.GetConnection()

	data := &factoid{Name: strings.ToLower(key)}

	res := db.Table("factoids").Where(data).FirstOrInit(data)
	if res.Error != nil {
		return nil
	}

	if data.Content == "" {
		return nil
	}

	parts := strings.Split(data.Content, ";;")
	cleaned := make([]string, 0)
	for _, v := range parts {
		if strings.TrimSpace(v) != "" {
			cleaned = append(cleaned, v)
		}
	}

	return cleaned
}

type factoid struct {
	Name    string `gorm:"name"`
	Content string `gorm:"content"`
}
