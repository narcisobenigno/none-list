package es

import (
	"errors"
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
	store    map[AggregateID]map[Version]StoredEvent
	position uint64
	mutex    sync.Mutex
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		store:    map[AggregateID]map[Version]StoredEvent{},
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
	sort.Slice(storedEvents, func(i, j int) bool {
		return storedEvents[i].Position < storedEvents[j].Position
	})

	return storedEvents, nil
}

func (i *InMemoryEventStore) Write(events []Event) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	storeCopy := i.storeCopy()

	for _, event := range events {
		if storeCopy[event.AggregateID()] == nil {
			storeCopy[event.AggregateID()] = map[Version]StoredEvent{}
		}
		if _, found := storeCopy[event.AggregateID()][event.AggregateVersion()]; found {
			return errors.New("optimistic lock violation")
		}

		storeCopy[event.AggregateID()][event.AggregateVersion()] = StoredEvent{
			Position: i.position,
			Event:    event,
		}
		i.position += 1
	}

	i.store = storeCopy

	return nil
}

func (i *InMemoryEventStore) storeCopy() map[AggregateID]map[Version]StoredEvent {
	storeCopy := map[AggregateID]map[Version]StoredEvent{}
	for aggregateID, versionEvent := range i.store {
		storeCopy[aggregateID] = map[Version]StoredEvent{}
		for version, event := range versionEvent {
			storeCopy[aggregateID][version] = event
		}
	}
	return storeCopy
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
