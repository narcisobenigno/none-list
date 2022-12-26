package list_test

import (
	"testing"

	"github.com/narcisobenigno/none-list/internal/domain/list"
	"github.com/narcisobenigno/none-list/pkg/es"
	"github.com/narcisobenigno/none-list/pkg/results"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CreateSuite struct {
	*require.Assertions
	suite.Suite
}

func TestProductSuite(t *testing.T) {
	s := &CreateSuite{}
	s.Assertions = require.New(t)
	suite.Run(t, s)
}

func (s *CreateSuite) TestCreatesList() {
	store := es.NewInMemoryEventStore()
	subject := list.NewBus(store)

	result, err := subject.Execute(&list.Create{
		ID:   es.NewDeterministicAggregateID("list-id-1"),
		Name: list.MustParseName("List name 1"),
	})
	s.NoError(err)

	s.Equal(results.Success(), result)
	s.Equal(
		[]es.StoredEvent{
			{
				Position: 1,
				Event: &list.Created{
					ID:   es.NewDeterministicAggregateID("list-id-1"),
					Name: list.MustParseName("List name 1"),
				},
			},
		},
		store.All(),
	)
}
