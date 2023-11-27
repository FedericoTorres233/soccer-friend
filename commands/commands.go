package commands

import (
	tele "gopkg.in/telebot.v3"
)

func Commands(b *tele.Bot) {
	firstCommand(b)
  get_teams(b)
  everything(b)
}
