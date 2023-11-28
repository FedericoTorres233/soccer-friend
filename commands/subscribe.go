package commands

import (
	"strings"

	tele "gopkg.in/telebot.v3"
)

func subscribe(b *tele.Bot) {
	userSubs := make(map[int64][]string)

	b.Handle("/subscribe", func(c tele.Context) error {

		chatID := c.Chat().ID
		arguments := c.Args()

		// Update map to assign a chat to 1 or more teams
		updatedUserSubs := append(userSubs[chatID], arguments...)
		userSubs[chatID] = updatedUserSubs

		// Check if arguments are nil
		if arguments == nil {
			return c.Send("Usage: /subscribe [team]")
		}

		// If successful send message to user
		c.Send("Subscribed to " + strings.Join(arguments, ", "))
		return c.Send("You are currently subscribed to:" + strings.Join(userSubs[chatID], ", "))
	})

}
