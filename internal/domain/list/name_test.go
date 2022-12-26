package list_test

import (
	"testing"

	"github.com/narcisobenigno/none-list/internal/domain/list"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type NameSuite struct {
	*require.Assertions
	suite.Suite
}

func TestNameSuite(t *testing.T) {
	s := new(NameSuite)
	s.Assertions = require.New(t)
	suite.Run(t, s)
}

func (s *NameSuite) TestParse() {
	s.Run("different names", func() {
		name, err := list.ParseName("List Name")
		s.NoError(err)

		anotherName, err := list.ParseName("Another list Name")
		s.NoError(err)

		s.NotEqual(name, anotherName)
	})

	s.Run("equal names", func() {
		name, err := list.ParseName("List Name")
		s.NoError(err)

		sameName, err := list.ParseName(" List Name ")
		s.NoError(err)

		s.Equal(name, sameName)
	})

	s.Run("returns error when empty name", func() {
		_, err := list.ParseName("  ")
		s.EqualError(err, "list name cannot be empty")
	})
}
