package list

import (
	"github.com/narcisobenigno/grocery-go/pkg/es"
	"github.com/narcisobenigno/grocery-go/pkg/results"
)

type Bus struct {
	store es.EventStore
}

func NewBus(store es.EventStore) *Bus {
	return &Bus{store: store}
}

func (b Bus) Execute(cmd es.Cmd) results.Result {
	list := newList()

	events, result := list.handle(cmd)

	b.store.Write(events)

	return result
}
