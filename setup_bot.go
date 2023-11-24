package main

import (
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func setupBot() (*tele.Bot, error) {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("Bot is up and running!")
	return b, nil
}
