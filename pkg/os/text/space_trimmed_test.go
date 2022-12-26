package text_test

import (
	"testing"

	"github.com/narcisobenigno/none-list/pkg/os/text"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SpaceTrimmedSuite struct {
	*require.Assertions
	suite.Suite
}

func TestSpaceTrimmedSuite(t *testing.T) {
	s := new(SpaceTrimmedSuite)
	s.Assertions = require.New(t)
	suite.Run(t, s)
}

func (s *SpaceTrimmedSuite) TestEquality() {
	spaceTrimmedText := text.NewSpaceTrimmed(" text with space trimmed")
	sameSpaceTrimmedText := text.NewSpaceTrimmed("text with space trimmed ")

	s.Equal(spaceTrimmedText, sameSpaceTrimmedText)

	spaceTrimmedText = text.NewSpaceTrimmed(" text with space trimmed")
	differentSpaceTrimmedText := text.NewSpaceTrimmed("different text with space trimmed ")

	s.NotEqual(spaceTrimmedText, differentSpaceTrimmedText)
}

func (s *SpaceTrimmedSuite) TestString() {
	s.Equal(
		"text with space trimmed",
		text.NewSpaceTrimmed(" text with space trimmed ").String(),
	)
}

func (s *SpaceTrimmedSuite) TestEmpty() {
	s.True(text.NewSpaceTrimmed("").Empty())
	s.True(text.NewSpaceTrimmed(" ").Empty())
	s.False(text.NewSpaceTrimmed(" a ").Empty())
}
