package commands

import (
	tele "gopkg.in/telebot.v3"
	"log"
)

func everything(b *tele.Bot) {
	b.Handle("/msg", func(c tele.Context) error {
		log.Println(c.Args())
		return nil
	})
}
