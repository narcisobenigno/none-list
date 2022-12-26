package results

import (
	"fmt"

	"github.com/narcisobenigno/none-list/pkg/os/texts"
)

type Result struct {
	fails map[texts.SpaceTrimmed]texts.SpaceTrimmed
}

func Success() Result {
	return newResult(map[texts.SpaceTrimmed]texts.SpaceTrimmed{})
}

func Failed(context, message string) Result {
	return newResult(map[texts.SpaceTrimmed]texts.SpaceTrimmed{
		texts.NewSpaceTrimmed(context): texts.NewSpaceTrimmed(message),
	})
}

func newResult(fails map[texts.SpaceTrimmed]texts.SpaceTrimmed) Result {
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
