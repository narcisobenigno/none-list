package list

import (
	"github.com/narcisobenigno/grocery-go/pkg/os/texts"
	"github.com/narcisobenigno/grocery-go/pkg/os/texts/spacetrimmed"
	"github.com/narcisobenigno/grocery-go/pkg/results"
)

type Name struct {
	name texts.Text
}

func TryParseName(name string) (Name, results.Result) {
	trimmedName := spacetrimmed.New(name)
	if trimmedName.Empty() {
		return Name{}, results.Failed("List", "name cannot be empty")
	}

	return Name{name: trimmedName}, results.Success()
}

func ParseName(name string) Name {
	parsed, result := TryParseName(name)
	if result.Failed() {
		panic(result.Message())
	}

	return parsed
}

func (n Name) Provided() bool {
	return !(n == Name{})
}
