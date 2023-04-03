package spacetrimmed_test

import (
	"testing"

	"github.com/narcisobenigno/grocery-go/pkg/os/texts/spacetrimmed"
	"github.com/stretchr/testify/require"
)

func TestEquality(t *testing.T) {
	spaceTrimmedText := spacetrimmed.New(" text with space trimmed")
	sameSpaceTrimmedText := spacetrimmed.New("text with space trimmed ")

	require.Equal(t, spaceTrimmedText, sameSpaceTrimmedText)

	spaceTrimmedText = spacetrimmed.New(" text with space trimmed")
	differentSpaceTrimmedText := spacetrimmed.New("different text with space trimmed ")

	require.NotEqual(t, spaceTrimmedText, differentSpaceTrimmedText)
}

func TestString(t *testing.T) {
	require.Equal(t,
		"text with space trimmed",
		spacetrimmed.New(" text with space trimmed ").String(),
	)
}

func TestEmpty(t *testing.T) {
	require.True(t, spacetrimmed.New("").Empty())
	require.True(t, spacetrimmed.New(" ").Empty())
	require.False(t, spacetrimmed.New(" a ").Empty())
}
