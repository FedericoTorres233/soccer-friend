package commands

import (
	tele "gopkg.in/telebot.v3"
)

var commands = [5]string{
	"start",
	"teams",
	"msg",
	"subscribe",
	"last",
}

func Commands(b *tele.Bot, API_KEY string) {
	firstCommand(b)
	get_teams(b, API_KEY)
	everything(b)
	subscribe(b)
	get_last_matches(b, API_KEY)
}
