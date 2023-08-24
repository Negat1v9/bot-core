package telegram

import (
	"fmt"
	"log"
	"strings"
)

const (
	cmdStart = "/start"
	cmdHelp  = "/help"
)

func (f *Fetcher) cmd(text string, userName string, ChatID int) error {
	text = strings.TrimSpace(text)
	log.Printf("new command %s from %s", text, userName)
	switch text {
	case cmdStart:
		return f.sayHello(ChatID, userName)
	case cmdHelp:
		return f.sayHelp(text, ChatID)
	default:
		return f.sayDefault(ChatID)
	}
}

func (f *Fetcher) sayHello(ChatID int, userName string) error {
	err := f.client.SendMessage(ChatID, fmt.Sprintf("%s %s", helloMsg, userName))
	if err != nil {
		return err
	}
	return nil
}

func (f *Fetcher) sayHelp(text string, ChatID int) error {
	if err := f.client.SendMessage(ChatID, helpMsg); err != nil {
		return err
	}
	return nil
}

func (f *Fetcher) sayDefault(ChatID int) error {
	if err := f.client.SendMessage(ChatID, defaultMsg); err != nil {
		return err
	}
	return nil
}
