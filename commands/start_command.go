package commands

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func startCommand(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		msg := fmt.Sprintf("Hi %v!\nType / to check all my available commands.\nContact me at @federicotorres", c.Chat().FirstName)
		return c.Send(msg)
	})
}
