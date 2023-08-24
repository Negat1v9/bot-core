package telegram

import (
	"errors"

	"github.com/Negat1v9/bot-core/bot/client"
	"github.com/Negat1v9/bot-core/bot/events"
)

type Fetcher struct {
	client *client.Client
	// storage
	offset int
}

// Add storage
func New(tg *client.Client) *Fetcher {
	return &Fetcher{
		client: tg,
	}
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown type of metadata")
)

// Meta for telegram
type Meta struct {
	UserName string
	ChatID   int
}

// fetcher for updates from server and convert updates in event type.
func (f *Fetcher) Fetch(limit int) ([]events.Event, error) {
	updates, err := f.client.Update(f.offset, limit)
	if err != nil {
		return nil, err
	}
	if len(updates) == 0 {
		return nil, nil
	}
	events := make([]events.Event, 0, len(updates))
	for _, upd := range updates {
		events = append(events, event(upd))
	}
	// update offset parametr
	f.offset = updates[len(updates)-1].Id + 1
	return events, nil
}

func (f *Fetcher) Direct(event events.Event) error {
	switch event.Type {
	case events.TypeMessage:
		return f.messageEvent(event)
	default:
		return ErrUnknownEventType
	}

}

func (f *Fetcher) messageEvent(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return err
	}
	if err := f.cmd(event.Text, meta.UserName, meta.ChatID); err != nil {
		return err
	}
	return nil
}

func meta(event events.Event) (Meta, error) {
	meta, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, ErrUnknownMetaType
	}
	return meta, nil
}

// Make from update to event type
func event(upd client.Updates) events.Event {
	updType := fetchType(upd)

	res := events.Event{
		Type: updType,
		Text: fethMessage(upd),
	}
	if updType == events.TypeMessage {
		res.Meta = Meta{
			UserName: upd.Message.From.UserName,
			ChatID:   upd.Message.Chat.Id,
		}
	}
	return res

}

// Catch type of event.
func fetchType(upd client.Updates) events.Type {
	if upd.Message == nil {
		return events.UnknownType
	}
	return events.TypeMessage
}

// Catch text of message from update
func fethMessage(upd client.Updates) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}
