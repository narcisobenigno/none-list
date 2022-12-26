package results_test

import (
	"testing"

	"github.com/narcisobenigno/none-list/pkg/results"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ResultSuite struct {
	*require.Assertions
	suite.Suite
}

func TestResultSuite(t *testing.T) {
	s := new(ResultSuite)
	s.Assertions = require.New(t)
	suite.Run(t, s)
}

func (s *ResultSuite) TestEquality() {
	s.Equal(results.Success(), results.Success())
	s.Equal(results.Failed("Context", "fail message"), results.Failed("Context", "fail message"))

	s.NotEqual(results.Success(), results.Failed("Context", "fail message"))
	s.NotEqual(results.Failed("Context", "fail message"), results.Failed("Context", "another fail message"))
	s.NotEqual(results.Failed("Context", "fail message"), results.Failed("AnotherContext", "fail message"))
}

func (s *ResultSuite) TestFailed() {
	s.False(results.Success().Failed())
	s.True(results.Failed("Context", "failed by").Failed())
}

func (s *ResultSuite) TestMessage() {
	s.Equal("", results.Success().Message())
	s.Equal("Context: fail message", results.Failed("Context", "fail message").Message())
	s.Equal("Context: fail message", results.Failed(" Context ", " fail message ").Message())
}
