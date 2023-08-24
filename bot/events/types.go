package events

type Type int

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

// inteface for give event to hadler
type Handler interface {
	Direct(e Event) error
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
