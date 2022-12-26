package text

import "strings"

type SpaceTrimmed struct {
	text string
}

func NewSpaceTrimmed(text string) SpaceTrimmed {
	return SpaceTrimmed{
		text: strings.TrimSpace(text),
	}
}

func (t SpaceTrimmed) String() string {
	return t.text
}
