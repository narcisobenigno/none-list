package es

type Event interface {
	AggregateID() AggregateID
	AggregateName() string
	AggregateVersion() uint64
}
