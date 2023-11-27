package commands

import (
	tele "gopkg.in/telebot.v3"
)

var commands = [4]string{
	"start",
	"teams",
	"msg",
	"subscribe",
}

func Commands(b *tele.Bot) {
	firstCommand(b)
	get_teams(b)
	everything(b)
	subscribe(b)
}
