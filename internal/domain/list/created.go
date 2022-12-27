package list

import (
	"github.com/narcisobenigno/none-list/pkg/es"
)

type Created struct {
	ID      es.AggregateID
	Name    Name
	Version es.Version
}

func (c Created) AggregateID() es.AggregateID {
	return c.ID
}

func (c Created) AggregateName() string {
	//TODO implement me
	panic("implement me")
}

func (c Created) AggregateVersion() es.Version {
	return c.Version
}
