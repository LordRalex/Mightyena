package mcping

import (
	"bytes"
	"fmt"
	"github.com/lordralex/mightyena/events"
	"github.com/lordralex/mightyena/services"
	"os/exec"
	"strings"
)

const ModuleName = "mcping"

func Load() {
	services.RegisterCommand(ModuleName, "mcping", runCommand)
}

func runCommand(e *events.Command) {
	if len(e.Arguments) != 1 {
		e.Respond("Usage: mcping <server>")
	}

	go execute(e, e.Arguments[0])
}

func execute(event *events.Command, addr string) {
	cmd := exec.Command("python", "mcping.py", addr)
	out, err := cmd.Output()
	if err != nil {
		event.Respond(fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	var buf bytes.Buffer
	buf.Write(out)
	event.Respond(strings.Replace(buf.String(), "\n", " - ", -1))
}
