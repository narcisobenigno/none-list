package list_test

import (
	"testing"

	"github.com/narcisobenigno/grocery-go/internal/domain/list"
	"github.com/narcisobenigno/grocery-go/pkg/es"
	"github.com/narcisobenigno/grocery-go/pkg/estest"
	"github.com/narcisobenigno/grocery-go/pkg/results"
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
	store := estest.NewInMemoryEventStore()
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
					ID:      es.NewDeterministicAggregateID("list-id-1"),
					Name:    list.MustParseName("List name 1"),
					Version: es.MustParseVersion(1),
				},
			},
		},
		store.All(),
	)
}

func (s *CreateSuite) TestFailsWhenNameNotProvided() {
	store := estest.NewInMemoryEventStore()
	subject := list.NewBus(store)

	result, err := subject.Execute(&list.Create{
		ID:   es.NewDeterministicAggregateID("list-id-1"),
		Name: list.Name{},
	})
	s.NoError(err)

	s.Equal(results.Failed("List", "name not provided"), result)
	s.Empty(store.All())
}
