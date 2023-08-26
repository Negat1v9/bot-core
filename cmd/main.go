package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/Negat1v9/bot-core/bot"
)

func main() {
	config := bot.NewConfig()
	_, err := toml.DecodeFile("config/bot-config.toml", config)
	if err != nil {
		fmt.Println("can't decode confit err: ", err.Error())
		return
	}
	if err := bot.Start(config); err != nil {
		fmt.Println("bot is stopped err:", err.Error())
	}
}
