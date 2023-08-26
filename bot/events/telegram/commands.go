package telegram

import (
	"context"
	"fmt"
	"log"
	"strings"
)

const (
	cmdStart       = "/start"
	cmdHelp        = "/help"
	cmdSetMyPeople = "/setchat"
)

func (f *Fetcher) cmd(ctx context.Context, text string, userName string, ChatID int) error {
	text = strings.TrimSpace(text)
	log.Printf("new command %s from %s", text, userName)
	switch text {
	case cmdStart:
		return f.sayHello(ctx, ChatID, userName)
	case cmdHelp:
		return f.sayHelp(ctx, text, ChatID)
	default:
		return f.sayDefault(ctx, ChatID)
	}
}

func (f *Fetcher) sayHello(ctx context.Context, ChatID int, userName string) error {
	err := f.client.SendMessage(ctx, ChatID, fmt.Sprintf("%s %s", helloMsg, userName))
	if err != nil {
		return err
	}
	return nil
}

func (f *Fetcher) sayHelp(ctx context.Context, text string, ChatID int) error {
	if err := f.client.SendMessage(ctx, ChatID, helpMsg); err != nil {
		return err
	}
	return nil
}

func (f *Fetcher) sayDefault(ctx context.Context, ChatID int) error {
	if err := f.client.SendMessage(ctx, ChatID, defaultMsg); err != nil {
		return err
	}
	return nil
}
