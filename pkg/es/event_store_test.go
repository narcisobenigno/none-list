package es_test

import (
	"testing"

	"github.com/narcisobenigno/none-list/pkg/es"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type EventStoreSuite struct {
	*require.Assertions
	suite.Suite
}

func TestEventStoreSuite(t *testing.T) {
	s := new(EventStoreSuite)
	s.Assertions = require.New(t)
	suite.Run(t, s)
}

func (s *EventStoreSuite) TestEventsAggregateByID() {
	store := es.NewInMemoryEventStore()

	err := store.Write([]es.Event{
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
	s.NoError(err)

	events, err := store.EventsByAggregateID(es.NewDeterministicAggregateID("something-happened-1"))
	s.NoError(err)
	s.Equal(
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

func (s *EventStoreSuite) TestInMemoryReturnsAllEvents() {
	store := es.NewInMemoryEventStore()

	err := store.Write([]es.Event{
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
	s.NoError(err)

	events := store.All()
	s.Equal(
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

func (s *EventStoreSuite) TestReturnsErrorWhenAggregateByID() {
	store := es.NewInMemoryEventStore()

	err := store.Write([]es.Event{
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
	s.NoError(err)

	err = store.Write([]es.Event{
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
	s.EqualError(err, "optimistic lock violation")

	events := store.All()
	s.Equal(
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
