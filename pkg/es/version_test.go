package es_test

import (
	"testing"

	"github.com/narcisobenigno/grocery-go/pkg/es"
	"github.com/narcisobenigno/grocery-go/pkg/results"
	"github.com/stretchr/testify/require"
)

func TestTryParseVersion(t *testing.T) {
	t.Run("equality fails when different versions", func(t *testing.T) {
		version, result := es.TryParseVersion(2)
		require.Equal(t, results.Success(), result)

		otherVersion, result := es.TryParseVersion(3)
		require.Equal(t, results.Success(), result)

		require.NotEqual(t, version, otherVersion)
	})

	t.Run("equality matches when same version", func(t *testing.T) {
		version, result := es.TryParseVersion(2)
		require.Equal(t, results.Success(), result)

		sameVersion, result := es.TryParseVersion(2)
		require.Equal(t, results.Success(), result)

		require.Equal(t, version, sameVersion)
	})

	t.Run("rejects when not acceptable version", func(t *testing.T) {
		_, result := es.TryParseVersion(0)
		require.Equal(t,
			results.Failed("Event", "version should be greater than or equal to 1"),
			result,
		)
	})
}

func TestMustParseVersion(t *testing.T) {
	t.Run("equality fails when different versions", func(t *testing.T) {
		require.NotEqual(t,
			es.ParseVersion(2),
			es.ParseVersion(3),
		)
	})

	t.Run("equality matches when same version", func(t *testing.T) {
		require.Equal(t,
			es.ParseVersion(2),
			es.ParseVersion(2),
		)
	})

	t.Run("panics when versions is smaller or equal zero", func(t *testing.T) {
		require.PanicsWithValue(t, "Event: version should be greater than or equal to 1", func() {
			es.ParseVersion(0)
		})
	})
}

func TestInitialVersion(t *testing.T) {
	require.Equal(t,
		es.ParseVersion(1),
		es.InitialVersion(),
	)
}
