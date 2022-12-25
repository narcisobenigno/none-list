package es

import "sync"

type EventStore interface {
	EventsByAggregateID(aggregateID AggregateID) ([]StoredEvent, error)
	Write(events []Event) error
}

type StoredEvent struct {
	Position uint64
	Event    Event
}

type InMemoryEventStore struct {
	store    map[AggregateID][]StoredEvent
	position uint64
	mutex    sync.Mutex
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		store:    map[AggregateID][]StoredEvent{},
		position: 1,
		mutex:    sync.Mutex{},
	}
}

func (i *InMemoryEventStore) EventsByAggregateID(aggregateID AggregateID) ([]StoredEvent, error) {
	return i.store[aggregateID], nil
}

func (i *InMemoryEventStore) Write(events []Event) error {
	i.mutex.Lock()
	for _, event := range events {
		i.store[event.AggregateID()] = append(i.store[event.AggregateID()], StoredEvent{
			Position: i.position,
			Event:    event,
		})
		i.position += 1
	}
	i.mutex.Unlock()
	return nil
}
