package es

import "github.com/google/uuid"

type AggregateID struct {
	uuid uuid.UUID
}

func NewDeterministicAggregateID(name string) AggregateID {
	return AggregateID{uuid: uuid.NewSHA1(uuid.NameSpaceOID, []byte(name))}
}
