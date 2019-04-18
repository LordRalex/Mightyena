package listeners

import (
	"github.com/lordralex/mightyena/modules/antispam"
	"github.com/lordralex/mightyena/modules/factoid"
)

func RegisterModules() {
	antispam.Load()

	factoid.Load()
}
