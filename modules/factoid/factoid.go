package factoid

import (
	"github.com/lordralex/mightyena/database"
	"github.com/lordralex/mightyena/events"
	"github.com/lordralex/mightyena/services"
	"regexp"
	"strings"
)

const ModuleName = "factoids"

var Bold = regexp.MustCompile(`\[/?b\]`)
var Underline = regexp.MustCompile(`\[/?u\]`)

func Load() {
	services.RegisterCommand(ModuleName, ">", globalHandle)
	services.RegisterCommand(ModuleName, "<", globalHandle)
	services.RegisterCommand(ModuleName, ".", globalHandle)
	services.RegisterCommand(ModuleName, "", globalHandle)
}

func globalHandle(command events.Command) {

}

func getFactoid(key string) []string {
	db := database.GetConnection()

	data := &factoid{Key: strings.ToLower(key)}

	res := db.Table("factoids").Where(data).FirstOrInit(data)
	if res.Error != nil {
		return make([]string, 0)
	}

	return strings.Split(data.Content, "|")
}

type factoid struct {
	Key     string `gorm:"name"`
	Content string `gorm:"content"`
}
