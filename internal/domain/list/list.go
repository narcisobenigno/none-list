package list

import (
	"github.com/narcisobenigno/grocery-go/pkg/es"
	"github.com/narcisobenigno/grocery-go/pkg/results"
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
	if !command.Name.Provided() {
		return []es.Event{}, results.Failed("List", "name not provided")
	}

	return []es.Event{
		&Created{
			ID:      command.ID,
			Name:    command.Name,
			Version: es.InitialVersion(),
		},
	}, results.Success()
}
