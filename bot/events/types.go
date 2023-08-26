package events

import "context"

type Type int

type Fetcher interface {
	Fetch(ctx context.Context, limit int) ([]Event, error)
}

// inteface for give event to hadler
type Handler interface {
	Direct(ctx context.Context, e Event) error
}

const (
	TypeMessage Type = iota
	UnknownType
)

type Event struct {
	Type Type
	Text string
	Meta interface{}
}
