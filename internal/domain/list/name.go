package list

import (
	"github.com/narcisobenigno/grocery-go/pkg/os/texts"
	"github.com/narcisobenigno/grocery-go/pkg/results"
)

type Name struct {
	name texts.Text
}

func ParseName(name string) (Name, results.Result) {
	trimmedName := texts.NewSpaceTrimmed(name)
	if trimmedName.Empty() {
		return Name{}, results.Failed("List", "name cannot be empty")
	}

	return Name{name: trimmedName}, results.Success()
}

func MustParseName(name string) Name {
	parsed, result := ParseName(name)
	if result.Failed() {
		panic(result.Message())
	}

	return parsed
}

func (n Name) Provided() bool {
	return !(n == Name{})
}
