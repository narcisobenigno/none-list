package texts_test

import (
	"testing"

	"github.com/narcisobenigno/none-list/pkg/os/texts"
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
	spaceTrimmedText := texts.NewSpaceTrimmed(" text with space trimmed")
	sameSpaceTrimmedText := texts.NewSpaceTrimmed("text with space trimmed ")

	s.Equal(spaceTrimmedText, sameSpaceTrimmedText)

	spaceTrimmedText = texts.NewSpaceTrimmed(" text with space trimmed")
	differentSpaceTrimmedText := texts.NewSpaceTrimmed("different text with space trimmed ")

	s.NotEqual(spaceTrimmedText, differentSpaceTrimmedText)
}

func (s *SpaceTrimmedSuite) TestString() {
	s.Equal(
		"text with space trimmed",
		texts.NewSpaceTrimmed(" text with space trimmed ").String(),
	)
}

func (s *SpaceTrimmedSuite) TestEmpty() {
	s.True(texts.NewSpaceTrimmed("").Empty())
	s.True(texts.NewSpaceTrimmed(" ").Empty())
	s.False(texts.NewSpaceTrimmed(" a ").Empty())
}
