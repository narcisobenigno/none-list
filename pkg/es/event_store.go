package es

type EventStore interface {
	EventsByAggregateID(aggregateID AggregateID) []StoredEvent
	Write(events []Event)
}

type StoredEvent struct {
	Position uint64
	Event    Event
}
