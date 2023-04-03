package list_test

import (
	"testing"

	"github.com/narcisobenigno/grocery-go/internal/domain/list"
	"github.com/narcisobenigno/grocery-go/pkg/es"
	"github.com/narcisobenigno/grocery-go/pkg/estest"
	"github.com/narcisobenigno/grocery-go/pkg/results"
	"github.com/stretchr/testify/require"
)

func TestCreatesList(t *testing.T) {
	store := estest.NewInMemoryEventStore()
	subject := list.NewBus(store)

	result := subject.Execute(&list.Create{
		ID:   es.NewDeterministicAggregateID("list-id-1"),
		Name: list.ParseName("List name 1"),
	})

	require.Equal(t, results.Success(), result)
	require.Equal(t,
		[]es.StoredEvent{
			{
				Position: 1,
				Event: &list.Created{
					ID:      es.NewDeterministicAggregateID("list-id-1"),
					Name:    list.ParseName("List name 1"),
					Version: es.ParseVersion(1),
				},
			},
		},
		store.All(),
	)
}

func TestFailsWhenNameNotProvided(t *testing.T) {
	store := estest.NewInMemoryEventStore()
	subject := list.NewBus(store)

	result := subject.Execute(&list.Create{
		ID:   es.NewDeterministicAggregateID("list-id-1"),
		Name: list.Name{},
	})

	require.Equal(t, results.Failed("List", "name not provided"), result)
	require.Empty(t, store.All())
}
