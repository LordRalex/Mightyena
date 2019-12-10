package services

import (
	"github.com/thoj/go-ircevent"
)

//var serviceLogger logging.Logger

func CreateServiceHandlers(connection *irc.Connection) {
	//service handlers
	connection.AddCallback("JOIN", handleJoin)
	connection.AddCallback("QUIT", handleQuit)
	connection.AddCallback("PART", handlePart)
	connection.AddCallback("KICK", handleKick)
	connection.AddCallback("NICK", handleNick)
	connection.AddCallback("353", handleNamesContinued)
	connection.AddCallback("366", handleNamesEnd)
	connection.AddCallback("352", handleWho)

	//event handler
	connection.AddCallback("PRIVMSG", fireMessageEvent)
	connection.AddCallback("JOIN", fireJoinEvent)
	connection.AddCallback("QUIT", fireQuitEvent)
	connection.AddCallback("PART", firePartEvent)
	connection.AddCallback("KICK", fireKickEvent)
	connection.AddCallback("NICK", fireNickEvent)
	connection.AddCallback("352", fireWhoEvent)

	//serviceLogger = logging.GetLogger("CORE SERVICE")

	startCleanupUserService()
}

func handleJoin(event *irc.Event) {
	handleJoinEventUserService(event)
}

func handleQuit(event *irc.Event) {
	handleQuitEventUserService(event)
	handleQuitEventChannelService(event)
}

func handlePart(event *irc.Event) {
	handlePartEventChannelService(event)
	handlePartEventUserService(event)
}

func handleNick(event *irc.Event) {
	handleNickEventUserService(event)
}

func handleNamesContinued(event *irc.Event) {
	handleNamesEventChannelService(event)
}

func handleNamesEnd(event *irc.Event) {
	handleNamesEndEventChannelService(event)
}

func handleKick(event *irc.Event) {
	//kicks are just forced parts to our services
	handlePartEventChannelService(event)
	handlePartEventUserService(event)
}

func handleWho(event *irc.Event) {
	handleWhoEventUserService(event)
}
