package main

import (
	"log"
	"os"

	"github.com/federicotorres233/soccer-friend/commands"
)

func main() {
	TG_TOKEN := os.Getenv("TG_TOKEN")

	b, err := setupBot(TG_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	commands.Commands(b)

	b.Start()
}
