package main

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Negat1v9/bot-core/bot"
	"github.com/Negat1v9/bot-core/bot/client"
	"github.com/Negat1v9/bot-core/bot/events/telegram"
	evlisten "github.com/Negat1v9/bot-core/bot/listener/event-listener"
)

const (
	tgHost        = "api.telegram.org"
	offsetUpdates = 50
)

func main() {
	config := bot.NewConfig()
	_, err := toml.DecodeFile("config/bot-config.toml", config)
	if err != nil {
		log.Fatalf("Error decode config file %s", err.Error())
	}
	// client for request to server
	client := client.New(tgHost, config.TgBotToken)
	// tgBot for work with telegram updates
	tgBot := telegram.New(client)
	// controller for listen updates and handle its
	controlListener := evlisten.New(tgBot, tgBot, 50)
	log.Println("start poolling bot")
	if err := controlListener.Start(); err != nil {
		log.Fatalf("server is stoped in %v", time.Now().Format(time.RFC1123))
	}
}
