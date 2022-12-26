package list

import (
	"github.com/narcisobenigno/none-list/pkg/es"
	"github.com/narcisobenigno/none-list/pkg/results"
)

type list struct {
}

func newList() list {
	return list{}
}

func (l list) handle(cmd es.Cmd) ([]es.Event, results.Result) {
	return create(cmd.(*Create))
}

func create(command *Create) ([]es.Event, results.Result) {
	return []es.Event{
		&Created{
			ID:   command.ID,
			Name: command.Name,
		},
	}, results.Success()
}
