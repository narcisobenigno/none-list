package es_test

import (
	"testing"

	"github.com/narcisobenigno/none-list/pkg/es"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AggregateIDSuite struct {
	*require.Assertions
	suite.Suite
}

func TestAggregateIDSuite(t *testing.T) {
	s := new(AggregateIDSuite)
	s.Assertions = require.New(t)
	suite.Run(t, s)
}

func (s *AggregateIDSuite) TestEquality() {
	s.NotEqual(
		es.NewDeterministicAggregateID("aggregate-id-1"),
		es.NewDeterministicAggregateID("aggregate-id-2"),
	)

	s.Equal(
		es.NewDeterministicAggregateID("aggregate-id-1"),
		es.NewDeterministicAggregateID("aggregate-id-1"),
	)
}
