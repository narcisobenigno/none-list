package list

import "github.com/narcisobenigno/grocery-go/pkg/es"

type AddItem struct {
	ListID es.AggregateID
	ID     string
}

func (a AddItem) AggregateName() string {
	return "LIST"
}

func (a AddItem) AggregateID() es.AggregateID {
	return a.ListID
}

func (a AddItem) Type() string {
	return "ADD_ITEM"
}
