package list_test

import (
	"testing"

	"github.com/narcisobenigno/none-list/internal/domain/list"
	"github.com/narcisobenigno/none-list/pkg/es"
	"github.com/narcisobenigno/none-list/pkg/results"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CreateListSuite struct {
	*require.Assertions
	suite.Suite
}

func TestProductSuite(t *testing.T) {
	s := &CreateListSuite{}
	s.Assertions = require.New(t)
	suite.Run(t, s)
}

func (s *CreateListSuite) TestCreatesList() {
	store := es.NewInMemoryEventStore()
	subject := list.NewBus(store)

	result := subject.Execute(&list.Create{
		ID:   es.NewDeterministicAggregateID("list-id-1"),
		Name: list.MustParseName("List name"),
	})

	s.Equal(results.Success(), result)
}
