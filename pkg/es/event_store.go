package es

type EventStore interface {
	EventsByAggregateID(aggregateID AggregateID) ([]StoredEvent, error)
	Write(events []Event) error
}

type StoredEvent struct {
	Position uint64
	Event    Event
}
