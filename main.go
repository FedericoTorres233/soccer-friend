package main

import (
	"github.com/federicotorres233/soccer-friend/commands"
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
