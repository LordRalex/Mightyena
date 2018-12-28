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
	//[SERVICE] [JOIN] &{Code:JOIN Raw::Mightyena!~Bot@beast.of.the.devilincarnate.net JOIN #minecrafthelp Nick:Mightyena Host:beast.of.the.devilincarnate.net Source:Mightyena!~Bot@beast.of.the.devilincarnate.net User:~Bot Arguments:[#minecrafthelp] Tags:map[] Connection:0xc0001aa000 Ctx:context.Background}
	//[SERVICE] [JOIN] &{Code:JOIN Raw::Lord_Ralex!~Ralex@irc2.ae97.net JOIN #minecrafthelp Nick:Lord_Ralex Host:irc2.ae97.net Source:Lord_Ralex!~Ralex@irc2.ae97.net User:~Ralex Arguments:[#minecrafthelp] Tags:map[] Connection:0xc0001aa000 Ctx:context.Background}
	fmt.Printf("[SERVICE] [JOIN] %+v\n", event)
}

func handleQuit(event *irc.Event) {
	fmt.Printf("[SERVICE] [QUIT] %+v\n", event)
}

func handlePart(event *irc.Event) {
	//[SERVICE] [PART] &{Code:PART Raw::Lord_Ralex!~Ralex@irc2.ae97.net PART #minecrafthelp :Leaving Nick:Lord_Ralex Host:irc2.ae97.net Source:Lord_Ralex!~Ralex@irc2.ae97.net User:~Ralex Arguments:[#minecrafthelp Leaving] Tags:map[] Connection:0xc0001aa000 Ctx:context.Background}
	fmt.Printf("[SERVICE] [PART] %+v\n", event)
}

func handleNick(event *irc.Event) {
	fmt.Printf("[SERVICE] [NICK] %+v\n", event)
}