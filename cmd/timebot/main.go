package main

import (
	"log"

	"github.com/kamaln7/timebot"
)

func main() {
	bot, err := timebot.New()
	if err != nil {
		log.Fatalln(err)
	}

	bot.Listen()
}
