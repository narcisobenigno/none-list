package results

import (
	"fmt"

	"github.com/narcisobenigno/grocery-go/pkg/os/texts/spacetrimmed"
)

type Result struct {
	fails map[spacetrimmed.SpaceTrimmed]spacetrimmed.SpaceTrimmed
}

func Success() Result {
	return newResult(map[spacetrimmed.SpaceTrimmed]spacetrimmed.SpaceTrimmed{})
}

func Failed(context, message string) Result {
	return newResult(map[spacetrimmed.SpaceTrimmed]spacetrimmed.SpaceTrimmed{
		spacetrimmed.New(context): spacetrimmed.New(message),
	})
}

func newResult(fails map[spacetrimmed.SpaceTrimmed]spacetrimmed.SpaceTrimmed) Result {
	return Result{fails: fails}
}

func (r Result) Failed() bool {
	return len(r.fails) > 0
}

func (r Result) Message() string {
	message := ""

	for context, failure := range r.fails {
		message = fmt.Sprintf("%s: %s", context, failure)
	}

	return message
}
