package bot

import (
	"database/sql"
	"log"

	"github.com/Negat1v9/bot-core/bot/client"
	"github.com/Negat1v9/bot-core/bot/events/telegram"
	evlisten "github.com/Negat1v9/bot-core/bot/listener/event-listener"
	"github.com/Negat1v9/bot-core/storage/sqlite"
)

const (
	tgHost        = "api.telegram.org"
	offsetUpdates = 50
)

func Start(conf *Config) error {

	client := client.New(tgHost, conf.TgBotToken)
	db, err := NewDB(conf.DBPath)
	defer db.Close()
	if err != nil {
		log.Fatalf("Can't create db. Bot is stoped - err: %s", err.Error())
	}
	// Storage for work with database
	storage := sqlite.New(db)
	// create table users.
	if err := storage.User().CreateTable(); err != nil {
		log.Fatalln("can't create table users err:", err.Error())
	}
	// tgBot for work with telegram updates
	tgBot := telegram.New(client, storage)
	// controller for listen updates and handle its
	controlListener := evlisten.New(tgBot, tgBot, 50)
	log.Println("start poolling bot")
	return controlListener.StartPooling()
}

func NewDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
