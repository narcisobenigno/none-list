package list

import "github.com/narcisobenigno/grocery-go/pkg/es"

type Create struct {
	ID   es.AggregateID
	Name Name
}

func (a Create) AggregateName() string {
	return "LIST"
}

func (a Create) AggregateID() es.AggregateID {
	return a.ID
}

func (a Create) Type() string {
	return "ADD_ITEM"
}
