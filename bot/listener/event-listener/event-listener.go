package evlisten

import (
	"context"
	"log"
	"time"

	"github.com/Negat1v9/bot-core/bot/events"
)

// type for cheking events from server
type Receiver struct {
	fetcher events.Fetcher
	handler events.Handler
	limit   int
}

func New(fetcher events.Fetcher, handler events.Handler, limit int) *Receiver {
	return &Receiver{
		fetcher: fetcher,
		handler: handler,
		limit:   limit,
	}
}

// Start Bot.
func (e *Receiver) StartPooling() error {
	for {
		evn, err := e.fetcher.Fetch(context.TODO(), e.limit)
		if err != nil {
			log.Printf("[Err]: receiver: %s", err.Error())
			continue
		}
		if len(evn) == 0 {
			time.Sleep(time.Second * 1)
			continue
		}
		// if not events wait
		if err := e.handleEvent(evn); err != nil {
			log.Print(err)
			continue
		}
	}
}

// Manage event.
func (e *Receiver) handleEvent(events []events.Event) error {
	start := time.Time{}
	for _, event := range events {
		start = time.Now()
		if err := e.handler.Direct(context.TODO(), event); err != nil {
			log.Printf("can't handle request: %s", err.Error())
			continue
		}
		log.Printf("time for response %v", time.Now().Sub(start))
	}
	return nil
}
