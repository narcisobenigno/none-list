package es

import (
	"sort"
	"sync"
)

type EventStore interface {
	EventsByAggregateID(aggregateID AggregateID) ([]StoredEvent, error)
	Write(events []Event) error
}

type StoredEvent struct {
	Position uint64
	Event    Event
}

type InMemoryEventStore struct {
	store    map[AggregateID]map[uint64]StoredEvent
	position uint64
	mutex    sync.Mutex
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		store:    map[AggregateID]map[uint64]StoredEvent{},
		position: 1,
		mutex:    sync.Mutex{},
	}
}

func (i *InMemoryEventStore) EventsByAggregateID(aggregateID AggregateID) ([]StoredEvent, error) {
	aggregateEvents := i.store[aggregateID]
	storedEvents := []StoredEvent{}
	for _, event := range aggregateEvents {
		storedEvents = append(storedEvents, event)
	}
	return storedEvents, nil
}

func (i *InMemoryEventStore) Write(events []Event) error {
	i.mutex.Lock()
	for _, event := range events {
		if i.store[event.AggregateID()] == nil {
			i.store[event.AggregateID()] = map[uint64]StoredEvent{}
		}
		i.store[event.AggregateID()][event.AggregateVersion()] = StoredEvent{
			Position: i.position,
			Event:    event,
		}
		i.position += 1
	}
	i.mutex.Unlock()
	return nil
}

func (i *InMemoryEventStore) All() []StoredEvent {
	events := []StoredEvent{}
	for _, storedEvents := range i.store {
		for _, event := range storedEvents {
			events = append(events, event)
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Position < events[j].Position
	})
	return events
}
