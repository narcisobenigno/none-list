package list_test

import (
	"testing"

	"github.com/narcisobenigno/grocery-go/internal/domain/list"
	"github.com/narcisobenigno/grocery-go/pkg/results"
	"github.com/stretchr/testify/require"
)

func TestTryParse(t *testing.T) {
	t.Run("different names", func(t *testing.T) {
		name, err := list.TryParseName("List Name")
		require.Equal(t, err, results.Success())

		anotherName, err := list.TryParseName("Another list Name")
		require.Equal(t, err, results.Success())

		require.NotEqual(t, name, anotherName)
	})

	t.Run("equal names", func(t *testing.T) {
		name, result := list.TryParseName("List Name")
		require.Equal(t, result, results.Success())

		sameName, result := list.TryParseName(" List Name ")
		require.Equal(t, result, results.Success())

		require.Equal(t, name, sameName)
	})

	t.Run("returns error when empty name", func(t *testing.T) {
		_, result := list.TryParseName("  ")
		require.Equal(t, results.Failed("List", "name cannot be empty"), result)
	})
}

func TestParseName(t *testing.T) {
	t.Run("different names", func(t *testing.T) {
		require.NotEqual(t,
			list.ParseName("List Name"),
			list.ParseName("Another list Name"),
		)
	})

	t.Run("equal names", func(t *testing.T) {
		require.Equal(t,
			list.ParseName("List Name"),
			list.ParseName(" List Name "),
		)
	})

	require.PanicsWithValue(t, "List: name cannot be empty", func() {
		list.ParseName("  ")
	})
}

func TestProvided(t *testing.T) {
	require.True(t, list.ParseName("List name").Provided())
	require.False(t, list.Name{}.Provided())
}
