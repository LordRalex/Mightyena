package listeners

import (
	"github.com/lordralex/mightyena/modules/antispam"
	"github.com/lordralex/mightyena/modules/factoid"
	"github.com/lordralex/mightyena/modules/webclient"
	"github.com/lordralex/mightyena/modules/welcome"
)

func RegisterModules() {
	antispam.Load()
	factoid.Load()
	webclient.Load()
	welcome.Load()
}
