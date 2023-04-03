package simple_test

import (
	"testing"

	"github.com/narcisobenigno/grocery-go/pkg/os/texts/simple"
	"github.com/stretchr/testify/require"
)

func TestEquality(t *testing.T) {
	t.Run("different texts", func(t *testing.T) {
		require.NotEqual(t,
			simple.New(" a text "),
			simple.New("a text"),
		)
	})
	t.Run("equal texts", func(t *testing.T) {
		require.Equal(t,
			simple.New("a text"),
			simple.New("a text"),
		)
	})
}

func TestString(t *testing.T) {
	require.Equal(t,
		" a text ",
		simple.New(" a text ").String(),
	)
}

func TestEmpty(t *testing.T) {
	require.True(t, simple.New("").Empty())
	require.False(t, simple.New(" ").Empty())
}
