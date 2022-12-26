package list_test

import (
	"testing"

	"github.com/narcisobenigno/none-list/internal/domain/list"
	"github.com/narcisobenigno/none-list/pkg/results"
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
		s.Equal(err, results.Success())

		anotherName, err := list.ParseName("Another list Name")
		s.Equal(err, results.Success())

		s.NotEqual(name, anotherName)
	})

	s.Run("equal names", func() {
		name, result := list.ParseName("List Name")
		s.Equal(result, results.Success())

		sameName, result := list.ParseName(" List Name ")
		s.Equal(result, results.Success())

		s.Equal(name, sameName)
	})

	s.Run("returns error when empty name", func() {
		_, result := list.ParseName("  ")
		s.Equal(results.Failed("List", "name cannot be empty"), result)
	})
}

func (s *NameSuite) TestMustParse() {
	s.Run("different names", func() {
		s.NotEqual(
			list.MustParseName("List Name"),
			list.MustParseName("Another list Name"),
		)
	})

	s.Run("equal names", func() {
		s.Equal(
			list.MustParseName("List Name"),
			list.MustParseName(" List Name "),
		)
	})

	s.PanicsWithValue("List: name cannot be empty", func() {
		list.MustParseName("  ")
	})
}

func (s *NameSuite) TestProvided() {
	s.True(list.MustParseName("List name").Provided())
	s.False(list.Name{}.Provided())
}
