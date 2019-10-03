package services

import (
	"github.com/lordralex/mightyena/events"
)

func RegisterAction(moduleName string, function func(evt *events.Action)) {
	register("action", moduleName, wrapAction(function))
}

func RegisterAuth(moduleName string, function func(evt *events.Auth)) {
	register("message", moduleName, wrapAuth(function))
}

func RegisterBan(moduleName string, function func(evt *events.Ban)) {
	register("message", moduleName, wrapBan(function))
}

func RegisterDeop(moduleName string, function func(evt *events.Deop)) {
	register("message", moduleName, wrapDeop(function))
}

func RegisterDevoice(moduleName string, function func(evt *events.Devoice)) {
	register("message", moduleName, wrapDevoice(function))
}

func RegisterInvite(moduleName string, function func(evt *events.Invite)) {
	register("message", moduleName, wrapInvite(function))
}

func RegisterJoin(moduleName string, function func(evt *events.Join)) {
	register("message", moduleName, wrapJoin(function))
}

func RegisterKick(moduleName string, function func(evt *events.Kick)) {
	register("message", moduleName, wrapKick(function))
}

func RegisterMessage(moduleName string, function func(evt *events.Message)) {
	register("message", moduleName, wrapMessage(function))
}

func RegisterNotice(moduleName string, function func(evt *events.Notice)) {
	register("message", moduleName, wrapNotice(function))
}

func RegisterOp(moduleName string, function func(evt *events.Op)) {
	register("message", moduleName, wrapOp(function))
}

func RegisterPart(moduleName string, function func(evt *events.Part)) {
	register("message", moduleName, wrapPart(function))
}

func RegisterQuit(moduleName string, function func(evt *events.Quit)) {
	register("message", moduleName, wrapQuit(function))
}

func RegisterUnban(moduleName string, function func(evt *events.Unban)) {
	register("message", moduleName, wrapUnban(function))
}

func RegisterVoice(moduleName string, function func(evt *events.Voice)) {
	register("message", moduleName, wrapVoice(function))
}