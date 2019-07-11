package modules

import (
	"github.com/lordralex/mightyena/modules/antispam"
	"github.com/lordralex/mightyena/modules/factoid"
	"github.com/lordralex/mightyena/modules/remove"
	"github.com/lordralex/mightyena/modules/welcome"
)

func Load() {
	antispam.Load()
	factoid.Load()
	//mcping.Load()
	remove.Load()
	//webclient.Load()
	welcome.Load()
}