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
	//connection.AddCallback("333", handleNamesStart)
	connection.AddCallback("353", handleNamesContinued)
	connection.AddCallback("366", handleNamesEnd)
	serviceLogger = logging.GetLogger("SERVICE")

	startCleanupUserService()
}

func handleJoin(event *irc.Event) {
	//[SERVICE] [JOIN] &{Code:JOIN Raw::Mightyena!~Bot@beast.of.the.devilincarnate.net JOIN #minecrafthelp Nick:Mightyena Host:beast.of.the.devilincarnate.net Source:Mightyena!~Bot@beast.of.the.devilincarnate.net User:~Bot Arguments:[#minecrafthelp] Tags:map[] Connection:0xc0001aa000 Ctx:context.Background}
	//[SERVICE] [JOIN] &{Code:JOIN Raw::Lord_Ralex!~Ralex@irc2.ae97.net JOIN #minecrafthelp Nick:Lord_Ralex Host:irc2.ae97.net Source:Lord_Ralex!~Ralex@irc2.ae97.net User:~Ralex Arguments:[#minecrafthelp] Tags:map[] Connection:0xc0001aa000 Ctx:context.Background}
	serviceLogger.Log(logging.Debug, "[JOIN] %+v", event)

	handleJoinEventUserService(event)
	//join already does a /names, so no need to do anything special for the bot user
}

func handleQuit(event *irc.Event) {
	//[SERVICE] [QUIT] &{Code:QUIT Raw::ismail!webchat@1.2.3.4.static.ttnet.com.tr QUIT :Quit: webchat.esper.net Nick:ismail Host:1.2.3.4.static.ttnet.com.tr Source:ismail!webchat@1.2.3.4.static.ttnet.com.tr User:webchat Arguments:[Quit: webchat.esper.net] Tags:map[] Connection:0xc0001a8000 Ctx:context.Background}
	serviceLogger.Log(logging.Debug, "[QUIT] %+v", event)

	handleQuitEventUserService(event)
	handleQuitEventChannelService(event)
}

func handlePart(event *irc.Event) {
	//[SERVICE] [PART] &{Code:PART Raw::Lord_Ralex!~Ralex@irc2.ae97.net PART #minecrafthelp :Leaving Nick:Lord_Ralex Host:irc2.ae97.net Source:Lord_Ralex!~Ralex@irc2.ae97.net User:~Ralex Arguments:[#minecrafthelp Leaving] Tags:map[] Connection:0xc0001aa000 Ctx:context.Background}
	serviceLogger.Log(logging.Debug, "[PART] %+v", event)

	handlePartEventChannelService(event)
	if event.Connection.GetNick() == event.Nick {
		//handle channel part event
	}
}

func handleNick(event *irc.Event) {
	//[SERVICE] [NICK] &{Code:NICK Raw::Guest39872!~webrosc@1.2.3.4.cable.virginm.net NICK :webrosc Nick:Guest39872 Host:1.2.3.4.cable.virginm.net Source:Guest39872!~webrosc@1.2.3.4.cable.virginm.net User:~webrosc Arguments:[webrosc] Tags:map[] Connection:0xc0001a8000 Ctx:context.Background}

	serviceLogger.Log(logging.Debug, "[NICK] %+v", event)
	handleNickEventUserService(event)
}

func handleNamesContinued(event *irc.Event) {
	//[SERVICE] [INFO] [NAMES] &{Code:353 Raw::nova.esper.net 353 Mightyena @ #minecrafthelp.breakroom :Mightyena Greymagic27 urielsalis Capfan67 redstonehelper TrueWolves lol768 GreyBot +GreyVulpine +phroa Extreme_ +clarjon1 md_5 +Andrio AuBot webrosc Extreme payonel rymate1234 +tyteen4a03 AG_Clinton +Kealper Odysseius mattym zkxs_was_taken Anna +Lord_Ralex yuken Marzenia GrygrFlzr +xales thgilfodrol +Morrolan Mustek +Wug Afootpluto osuka +Absol Nick: Host: Source:nova.esper.net User: Arguments:[ Mightyena Greymagic27 urielsalis Capfan67 redstonehelper TrueWolves lol768 GreyBot +GreyVulpine +phroa Extreme_ +clarjon1 md_5 +Andrio AuBot webrosc Extreme payonel rymate1234 +tyteen4a03 AG_Clinton +Kealper Odysseius mattym zkxs_was_taken Anna +Lord_Ralex yuken Marzenia GrygrFlzr +xales thgilfodrol +Morrolan Mustek +Wug Afootpluto osuka +Absol] Tags:map[] Connection:0xc0001a8000 Ctx:context.Background}
	//discard first 2 args
	//third arg is the channel
	//rest are users
	handleNamesEventChannelService(event)
}

func handleNamesEnd(event *irc.Event) {
	//[SERVICE] [INFO] [NAMES-END] &{Code:366 Raw::nova.esper.net 366 Mightyena #minecrafthelp.breakroom :End of /NAMES list. Nick: Host: Source:nova.esper.net User: Arguments:[Mightyena #minecrafthelp.breakroom End of /NAMES list.] Tags:map[] Connection:0xc0001a8000 Ctx:context.Background}
	handleNamesEndEventChannelService(event)
}
