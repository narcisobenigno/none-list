package list

import (
	"errors"

	"github.com/narcisobenigno/none-list/pkg/os/texts"
)

type Name struct {
	name texts.Text
}

func ParseName(name string) (Name, error) {
	trimmedName := texts.NewSpaceTrimmed(name)
	if trimmedName.Empty() {
		return Name{}, errors.New("list name cannot be empty")
	}

	return Name{name: trimmedName}, nil
}

func MustParseName(name string) Name {
	return Name{}
}
