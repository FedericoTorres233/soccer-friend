package main

import (
	"log"

	"github.com/federicotorres233/soccer-friend/commands"
)

func main() {

	b, err := setupBot()
	if err != nil {
		log.Fatal(err)
	}

	commands.Commands(b)

	b.Start()
}
