package es

type Cmd interface {
	AggregateName() string
	AggregateID() AggregateID
	Type() string
}
