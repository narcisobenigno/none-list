package assert_test

import (
	"errors"
	"testing"

	"github.com/narcisobenigno/grocery-go/pkg/assert"
	"github.com/stretchr/testify/require"
)

func TestNoError(t *testing.T) {
	t.Run("panics when error present", func(t *testing.T) {
		require.PanicsWithError(t,
			"something wrong",
			func() {
				assert.NoError(errors.New("something wrong"))
			},
		)
	})

	t.Run("does not panics when no error", func(t *testing.T) {
		require.NotPanics(t,
			func() {
				assert.NoError(nil)
			},
		)
	})
}

func TestMust(t *testing.T) {
	t.Run("panics when error present", func(t *testing.T) {
		require.PanicsWithError(t,
			"something wrong",
			func() {
				assert.Must("I won't return", errors.New("something wrong"))
			},
		)
	})

	t.Run("does not panics when no error", func(t *testing.T) {
		require.NotPanics(t,
			func() {
				require.Equal(t, "I'll return", assert.Must("I'll return", nil))
			},
		)
	})
}

func TestTrue(t *testing.T) {
	t.Run("panics when false", func(t *testing.T) {
		require.PanicsWithValue(t,
			"ops false",
			func() {
				assert.True(false, "ops false")
			},
		)
	})

	t.Run("does not panic when true", func(t *testing.T) {
		require.NotPanics(t,
			func() {
				assert.True(true, "no panic")
			},
		)
	})
}

func TestFalse(t *testing.T) {
	t.Run("panics when true", func(t *testing.T) {
		require.PanicsWithValue(t,
			"ops true",
			func() {
				assert.False(true, "ops true")
			},
		)
	})

	t.Run("does not panic when false", func(t *testing.T) {
		require.NotPanics(t,
			func() {
				assert.False(false, "no panic")
			},
		)
	})
}
