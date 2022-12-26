package list

import (
	"github.com/narcisobenigno/none-list/pkg/es"
	"github.com/narcisobenigno/none-list/pkg/results"
)

type Bus struct {
	store es.EventStore
}

func NewBus(store es.EventStore) *Bus {
	return &Bus{store: store}
}

func (b Bus) Execute(cmd es.Cmd) (results.Result, error) {
	list := newList()

	events, result := list.handle(cmd)

	err := b.store.Write(events)
	if err != nil {
		panic(err)
	}

	return result, nil
}
