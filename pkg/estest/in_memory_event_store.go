package estest

import (
	"sort"
	"sync"

	"github.com/narcisobenigno/grocery-go/pkg/es"
)

type InMemoryEventStore struct {
	store    map[es.AggregateID]map[es.Version]es.StoredEvent
	position uint64
	mutex    sync.Mutex
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		store:    map[es.AggregateID]map[es.Version]es.StoredEvent{},
		position: 1,
		mutex:    sync.Mutex{},
	}
}

func (i *InMemoryEventStore) EventsByAggregateID(aggregateID es.AggregateID) []es.StoredEvent {
	aggregateEvents := i.store[aggregateID]
	storedEvents := []es.StoredEvent{}
	for _, event := range aggregateEvents {
		storedEvents = append(storedEvents, event)
	}
	sort.Slice(storedEvents, func(i, j int) bool {
		return storedEvents[i].Position < storedEvents[j].Position
	})

	return storedEvents
}

func (i *InMemoryEventStore) Write(events []es.Event) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	storeCopy := i.storeCopy()

	for _, event := range events {
		if storeCopy[event.AggregateID()] == nil {
			storeCopy[event.AggregateID()] = map[es.Version]es.StoredEvent{}
		}
		if _, found := storeCopy[event.AggregateID()][event.AggregateVersion()]; found {
			panic("optimistic lock violation")
		}

		storeCopy[event.AggregateID()][event.AggregateVersion()] = es.StoredEvent{
			Position: i.position,
			Event:    event,
		}
		i.position += 1
	}

	i.store = storeCopy
}

func (i *InMemoryEventStore) storeCopy() map[es.AggregateID]map[es.Version]es.StoredEvent {
	storeCopy := map[es.AggregateID]map[es.Version]es.StoredEvent{}
	for aggregateID, versionEvent := range i.store {
		storeCopy[aggregateID] = map[es.Version]es.StoredEvent{}
		for version, event := range versionEvent {
			storeCopy[aggregateID][version] = event
		}
	}
	return storeCopy
}

func (i *InMemoryEventStore) All() []es.StoredEvent {
	events := []es.StoredEvent{}
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
