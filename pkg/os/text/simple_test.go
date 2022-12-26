package text_test

import (
	"testing"

	"github.com/narcisobenigno/none-list/pkg/os/text"
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
			text.New(" a text "),
			text.New("a text"),
		)
	})
	s.Run("equal texts", func() {
		s.Equal(
			text.New("a text"),
			text.New("a text"),
		)
	})
}

func (s *SimpleSuite) TestString() {
	s.Equal(
		" a text ",
		text.New(" a text ").String(),
	)
}

func (s *SimpleSuite) TestEmpty() {
	s.True(text.New("").Empty())
	s.False(text.New(" ").Empty())
}
