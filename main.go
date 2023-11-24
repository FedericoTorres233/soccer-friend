package main

import (
	"github.com/federicotorres233/telebot/commands"
	"log"
)

func main() {

	b, err := setupBot()
	if err != nil {
		log.Fatal(err)
	}

	commands.Commands(b)

	b.Start()
}
