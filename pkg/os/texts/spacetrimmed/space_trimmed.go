package spacetrimmed

import (
	"strings"

	"github.com/narcisobenigno/grocery-go/pkg/os/texts"
	"github.com/narcisobenigno/grocery-go/pkg/os/texts/simple"
)

type SpaceTrimmed struct {
	text texts.Text
}

func New(text string) SpaceTrimmed {
	return SpaceTrimmed{
		text: simple.New(strings.TrimSpace(text)),
	}
}

func (t SpaceTrimmed) String() string {
	return t.text.String()
}

func (t SpaceTrimmed) Empty() bool {
	return t.text.Empty()
}
