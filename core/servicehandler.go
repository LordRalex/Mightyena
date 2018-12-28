package core

import (
	"fmt"
	"github.com/thoj/go-ircevent"
)

func CreateServiceHandlers(connection *irc.Connection) {
	connection.AddCallback("JOIN", handleJoin)
	connection.AddCallback("QUIT", handleQuit)
	connection.AddCallback("PART", handlePart)
	connection.AddCallback("NICK", handleNick)
}

func handleJoin(event *irc.Event) {
	fmt.Printf("[SERVICE] [JOIN] %+v\n", event)
}

func handleQuit(event *irc.Event) {
	fmt.Printf("[SERVICE] [QUIT] %+v\n", event)
}

func handlePart(event *irc.Event) {
	fmt.Printf("[SERVICE] [PART] %+v\n", event)
}

func handleNick(event *irc.Event) {
	fmt.Printf("[SERVICE] [NICK] %+v\n", event)
}