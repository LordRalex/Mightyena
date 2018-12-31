package core

import (
	"github.com/lordralex/mightyena/logging"
	"github.com/thoj/go-ircevent"
)

var serviceLogger logging.Logger

func CreateServiceHandlers(connection *irc.Connection) {
	connection.AddCallback("JOIN", handleJoin)
	connection.AddCallback("QUIT", handleQuit)
	connection.AddCallback("PART", handlePart)
	connection.AddCallback("NICK", handleNick)
	serviceLogger = logging.GetLogger("SERVICE")

	startCleanupUserService()
}

func handleJoin(event *irc.Event) {
	//[SERVICE] [JOIN] &{Code:JOIN Raw::Mightyena!~Bot@beast.of.the.devilincarnate.net JOIN #minecrafthelp Nick:Mightyena Host:beast.of.the.devilincarnate.net Source:Mightyena!~Bot@beast.of.the.devilincarnate.net User:~Bot Arguments:[#minecrafthelp] Tags:map[] Connection:0xc0001aa000 Ctx:context.Background}
	//[SERVICE] [JOIN] &{Code:JOIN Raw::Lord_Ralex!~Ralex@irc2.ae97.net JOIN #minecrafthelp Nick:Lord_Ralex Host:irc2.ae97.net Source:Lord_Ralex!~Ralex@irc2.ae97.net User:~Ralex Arguments:[#minecrafthelp] Tags:map[] Connection:0xc0001aa000 Ctx:context.Background}
	serviceLogger.Log(logging.Info, "[JOIN] %+v\n", event)

	handleJoinEventUserService(event)
	if event.Connection.GetNick() == event.Nick {
		//handle channel join event
	}
}

func handleQuit(event *irc.Event) {
	//[SERVICE] [QUIT] &{Code:QUIT Raw::ismail!webchat@1.2.3.4.static.ttnet.com.tr QUIT :Quit: webchat.esper.net Nick:ismail Host:1.2.3.4.static.ttnet.com.tr Source:ismail!webchat@1.2.3.4.static.ttnet.com.tr User:webchat Arguments:[Quit: webchat.esper.net] Tags:map[] Connection:0xc0001a8000 Ctx:context.Background}
	serviceLogger.Log(logging.Info, "[QUIT] %+v\n", event)

	handleQuitEventUserService(event)
	if event.Connection.GetNick() == event.Nick {
		//handle channel quit event
	}
}

func handlePart(event *irc.Event) {
	//[SERVICE] [PART] &{Code:PART Raw::Lord_Ralex!~Ralex@irc2.ae97.net PART #minecrafthelp :Leaving Nick:Lord_Ralex Host:irc2.ae97.net Source:Lord_Ralex!~Ralex@irc2.ae97.net User:~Ralex Arguments:[#minecrafthelp Leaving] Tags:map[] Connection:0xc0001aa000 Ctx:context.Background}
	serviceLogger.Log(logging.Info, "[PART] %+v\n", event)

	//handleJoinEventUserService(event)
	if event.Connection.GetNick() == event.Nick {
		//handle channel part event
	}
}

func handleNick(event *irc.Event) {
	//[SERVICE] [NICK] &{Code:NICK Raw::Guest39872!~webrosc@1.2.3.4.cable.virginm.net NICK :webrosc Nick:Guest39872 Host:1.2.3.4.cable.virginm.net Source:Guest39872!~webrosc@1.2.3.4.cable.virginm.net User:~webrosc Arguments:[webrosc] Tags:map[] Connection:0xc0001a8000 Ctx:context.Background}

	serviceLogger.Log(logging.Info, "[NICK] %+v\n", event)
	handleNickEventUserService(event)
}
