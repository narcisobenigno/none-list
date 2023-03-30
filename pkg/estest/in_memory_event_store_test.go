package estest_test

import (
	"testing"

	"github.com/narcisobenigno/grocery-go/pkg/es"
	"github.com/narcisobenigno/grocery-go/pkg/estest"
	"github.com/stretchr/testify/require"
)

func TestEventsAggregateByID(t *testing.T) {
	store := estest.NewInMemoryEventStore()
	store.Write([]es.Event{
		&somethingHappened{
			DataID:  es.NewDeterministicAggregateID("something-happened-1"),
			What:    "something happened",
			Version: es.MustParseVersion(1),
		},
		&somethingHappened{
			DataID:  es.NewDeterministicAggregateID("something-happened-2"),
			What:    "something happened 2",
			Version: es.MustParseVersion(1),
		},
		&somethingHappened{
			DataID:  es.NewDeterministicAggregateID("something-happened-1"),
			What:    "something happened again",
			Version: es.MustParseVersion(2),
		},
	})

	events := store.EventsByAggregateID(es.NewDeterministicAggregateID("something-happened-1"))
	require.Equal(t,
		[]es.StoredEvent{
			{
				Position: 1,
				Event: &somethingHappened{
					DataID:  es.NewDeterministicAggregateID("something-happened-1"),
					What:    "something happened",
					Version: es.MustParseVersion(1),
				},
			},
			{
				Position: 3,
				Event: &somethingHappened{
					DataID:  es.NewDeterministicAggregateID("something-happened-1"),
					What:    "something happened again",
					Version: es.MustParseVersion(2),
				},
			},
		},
		events,
	)
}

func TestInMemoryReturnsAllEvents(t *testing.T) {
	store := estest.NewInMemoryEventStore()
	store.Write([]es.Event{
		&somethingHappened{
			DataID:  es.NewDeterministicAggregateID("something-happened-1"),
			What:    "something happened",
			Version: es.MustParseVersion(1),
		},
		&somethingHappened{
			DataID:  es.NewDeterministicAggregateID("something-happened-2"),
			What:    "something happened 2",
			Version: es.MustParseVersion(1),
		},
		&somethingHappened{
			DataID:  es.NewDeterministicAggregateID("something-happened-1"),
			What:    "something happened again",
			Version: es.MustParseVersion(2),
		},
	})

	events := store.All()
	require.Equal(t,
		[]es.StoredEvent{
			{
				Position: 1,
				Event: &somethingHappened{
					DataID:  es.NewDeterministicAggregateID("something-happened-1"),
					What:    "something happened",
					Version: es.MustParseVersion(1),
				},
			},
			{
				Position: 2,
				Event: &somethingHappened{
					DataID:  es.NewDeterministicAggregateID("something-happened-2"),
					What:    "something happened 2",
					Version: es.MustParseVersion(1),
				},
			},
			{
				Position: 3,
				Event: &somethingHappened{
					DataID:  es.NewDeterministicAggregateID("something-happened-1"),
					What:    "something happened again",
					Version: es.MustParseVersion(2),
				},
			},
		},
		events,
	)
}

func TestReturnsErrorWhenAggregateByID(t *testing.T) {
	store := estest.NewInMemoryEventStore()
	store.Write([]es.Event{
		&somethingHappened{
			DataID:  es.NewDeterministicAggregateID("something-happened-1"),
			What:    "something happened",
			Version: es.MustParseVersion(1),
		},
		&somethingHappened{
			DataID:  es.NewDeterministicAggregateID("something-happened-2"),
			What:    "something happened 2",
			Version: es.MustParseVersion(1),
		},
		&somethingHappened{
			DataID:  es.NewDeterministicAggregateID("something-happened-1"),
			What:    "something happened again",
			Version: es.MustParseVersion(2),
		},
	})

	require.PanicsWithValue(t, "optimistic lock violation", func() {
		store.Write([]es.Event{
			&somethingHappened{
				DataID:  es.NewDeterministicAggregateID("something-happened-1"),
				What:    "existing version",
				Version: es.MustParseVersion(1),
			},
			&somethingHappened{
				DataID:  es.NewDeterministicAggregateID("something-happened-1"),
				What:    "non existing, but should be ignored anyway since the other event failed",
				Version: es.MustParseVersion(3),
			},
		})
	})

	events := store.All()
	require.Equal(t,
		[]es.StoredEvent{
			{
				Position: 1,
				Event: &somethingHappened{
					DataID:  es.NewDeterministicAggregateID("something-happened-1"),
					What:    "something happened",
					Version: es.MustParseVersion(1),
				},
			},
			{
				Position: 2,
				Event: &somethingHappened{
					DataID:  es.NewDeterministicAggregateID("something-happened-2"),
					What:    "something happened 2",
					Version: es.MustParseVersion(1),
				},
			},
			{
				Position: 3,
				Event: &somethingHappened{
					DataID:  es.NewDeterministicAggregateID("something-happened-1"),
					What:    "something happened again",
					Version: es.MustParseVersion(2),
				},
			},
		},
		events,
	)
}

type somethingHappened struct {
	DataID  es.AggregateID
	What    string
	Version es.Version
}

func (s somethingHappened) AggregateID() es.AggregateID {
	return s.DataID
}

func (s somethingHappened) AggregateName() string {
	return "SOMETHING"
}

func (s somethingHappened) AggregateVersion() es.Version {
	return s.Version
}
