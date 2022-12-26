package text

import "strings"

type SpaceTrimmed struct {
	text Text
}

func NewSpaceTrimmed(text string) SpaceTrimmed {
	return SpaceTrimmed{
		text: New(strings.TrimSpace(text)),
	}
}

func (t SpaceTrimmed) String() string {
	return t.text.String()
}

func (t SpaceTrimmed) Empty() bool {
	return t.text.Empty()
}
