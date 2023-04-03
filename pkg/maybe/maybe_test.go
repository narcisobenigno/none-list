package maybe_test

import (
	"testing"

	"github.com/narcisobenigno/grocery-go/pkg/maybe"
	"github.com/stretchr/testify/require"
)

func TestApply(t *testing.T) {
	sum := func(x, y int) int { return x + y }
	thirty := 30
	twenty := 20
	require.Equal(t, maybe.Some(50), maybe.Apply(
		maybe.ApplyFunctor(
			maybe.Some[func(int, int) int](sum),
			maybe.Some(thirty),
		),
		maybe.Some(twenty),
	))

}
