package es_test

import (
	"testing"

	"github.com/narcisobenigno/grocery-go/pkg/es"
	"github.com/stretchr/testify/require"
)

func TestEquality(t *testing.T) {
	require.NotEqual(t,
		es.NewDeterministicAggregateID("aggregate-id-1"),
		es.NewDeterministicAggregateID("aggregate-id-2"),
	)

	require.Equal(t,
		es.NewDeterministicAggregateID("aggregate-id-1"),
		es.NewDeterministicAggregateID("aggregate-id-1"),
	)
}
