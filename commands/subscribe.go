package commands

import (
	tele "gopkg.in/telebot.v3"
	"log"
)

func subscribe(b *tele.Bot) {
	b.Handle("/subscribe", func(c tele.Context) error {
		log.Println(c.Args())
		if c.Args() == nil {
			return c.Send("Usage: /subscribe [team]")
		}
		return c.Send("Subscribed to " + c.Args()[0])
	})
}
