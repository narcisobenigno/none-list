package texts_test

import (
	"testing"

	"github.com/narcisobenigno/none-list/pkg/os/texts"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SimpleSuite struct {
	*require.Assertions
	suite.Suite
}

func TestSimpleSuite(t *testing.T) {
	s := new(SimpleSuite)
	s.Assertions = require.New(t)
	suite.Run(t, s)
}

func (s *SimpleSuite) TestEquality() {
	s.Run("different texts", func() {
		s.NotEqual(
			texts.New(" a text "),
			texts.New("a text"),
		)
	})
	s.Run("equal texts", func() {
		s.Equal(
			texts.New("a text"),
			texts.New("a text"),
		)
	})
}

func (s *SimpleSuite) TestString() {
	s.Equal(
		" a text ",
		texts.New(" a text ").String(),
	)
}

func (s *SimpleSuite) TestEmpty() {
	s.True(texts.New("").Empty())
	s.False(texts.New(" ").Empty())
}
