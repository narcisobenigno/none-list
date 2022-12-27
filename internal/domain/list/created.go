package list

import (
	"github.com/narcisobenigno/none-list/pkg/es"
)

type Created struct {
	ID      es.AggregateID
	Name    Name
	Version uint64
}

func (c Created) AggregateID() es.AggregateID {
	return c.ID
}

func (c Created) AggregateName() string {
	//TODO implement me
	panic("implement me")
}

func (c Created) AggregateVersion() uint64 {
	return c.Version
}
