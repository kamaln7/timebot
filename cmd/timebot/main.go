package main

import (
	"github.com/kamaln7/timebot"
	"github.com/kamaln7/timebot/config"
)

func main() {
	conf := config.Read()

	bot := timebot.New(&timebot.Config{
		Host:      conf.Host,
		Timezones: conf.Timezones,
	})
	bot.Listen()
}
