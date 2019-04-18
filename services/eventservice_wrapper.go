package services

import "github.com/lordralex/mightyena/events"

func wrapAction(function func(message events.Action)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Action))
	}
}

func wrapAuth(function func(message events.Auth)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Auth))
	}
}

func wrapBan(function func(message events.Ban)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Ban))
	}
}

func wrapCommand(function func(message events.Command)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Command))
	}
}

func wrapDeop(function func(message events.Deop)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Deop))
	}
}

func wrapDevoice(function func(message events.Devoice)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Devoice))
	}
}

func wrapInvite(function func(message events.Invite)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Invite))
	}
}

func wrapJoin(function func(message events.Join)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Join))
	}
}

func wrapKick(function func(message events.Kick)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Kick))
	}
}

func wrapMessage(function func(message events.Message)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Message))
	}
}

func wrapNotice(function func(message events.Notice)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Notice))
	}
}

func wrapOp(function func(message events.Op)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Op))
	}
}

func wrapPart(function func(message events.Part)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Part))
	}
}

func wrapUnban(function func(message events.Unban)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Unban))
	}
}

func wrapVoice(function func(message events.Voice)) func(events.Event){
	return func(event events.Event) {
		function(event.(events.Voice))
	}
}

func wrapQuit(function func(message events.Quit)) func(event events.Event) {
	return func(event events.Event) {
		function(event.(events.Quit))
	}
}