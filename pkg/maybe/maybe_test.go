package maybe_test

import (
	"testing"

	"github.com/narcisobenigno/grocery-go/pkg/maybe"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MaybeSuite struct {
	*require.Assertions
	suite.Suite
}

func TestMaybeSuite(t *testing.T) {
	s := new(MaybeSuite)
	s.Assertions = require.New(t)
	suite.Run(t, s)
}

func (s *MaybeSuite) TestApply() {
	sum := func(x, y int) int { return x + y }
	thirty := 30
	twenty := 20
	s.Equal(maybe.Some(50), maybe.Apply(
		maybe.ApplyFunctor(
			maybe.Some[func(int, int) int](sum),
			maybe.Some(thirty),
		),
		maybe.Some(twenty),
	))

}
