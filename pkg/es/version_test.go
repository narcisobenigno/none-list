package es_test

import (
	"testing"

	"github.com/narcisobenigno/grocery-go/pkg/es"
	"github.com/narcisobenigno/grocery-go/pkg/results"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type VersionSuite struct {
	*require.Assertions
	suite.Suite
}

func TestVersionSuite(t *testing.T) {
	s := new(VersionSuite)
	s.Assertions = require.New(t)
	suite.Run(t, s)
}

func (s *VersionSuite) TestParse() {
	s.Run("equality fails when different versions", func() {
		version, result := es.ParseVersion(2)
		s.Equal(results.Success(), result)

		otherVersion, result := es.ParseVersion(3)
		s.Equal(results.Success(), result)

		s.NotEqual(version, otherVersion)
	})

	s.Run("equality matches when same version", func() {
		version, result := es.ParseVersion(2)
		s.Equal(results.Success(), result)

		sameVersion, result := es.ParseVersion(2)
		s.Equal(results.Success(), result)

		s.Equal(version, sameVersion)
	})

	s.Run("rejects when not acceptable version", func() {
		_, result := es.ParseVersion(0)
		s.Equal(
			results.Failed("Event", "version should be greater than or equal to 1"),
			result,
		)
	})
}

func (s *VersionSuite) TestMustParse() {
	s.Run("equality fails when different versions", func() {
		s.NotEqual(
			es.MustParseVersion(2),
			es.MustParseVersion(3),
		)
	})

	s.Run("equality matches when same version", func() {
		s.Equal(
			es.MustParseVersion(2),
			es.MustParseVersion(2),
		)
	})

	s.PanicsWithValue("Event: version should be greater than or equal to 1", func() {
		es.MustParseVersion(0)
	})
}

func (s *VersionSuite) TestInitialVersion() {
	s.Equal(
		es.MustParseVersion(1),
		es.InitialVersion(),
	)
}
