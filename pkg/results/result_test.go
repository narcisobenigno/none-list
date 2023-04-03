package results_test

import (
	"testing"

	"github.com/narcisobenigno/grocery-go/pkg/results"
	"github.com/stretchr/testify/require"
)

func TestEquality(t *testing.T) {
	require.Equal(t, results.Success(), results.Success())
	require.Equal(t, results.Failed("Context", "fail message"), results.Failed("Context", "fail message"))

	require.NotEqual(t, results.Success(), results.Failed("Context", "fail message"))
	require.NotEqual(t, results.Failed("Context", "fail message"), results.Failed("Context", "another fail message"))
	require.NotEqual(t, results.Failed("Context", "fail message"), results.Failed("AnotherContext", "fail message"))
}

func TestFailed(t *testing.T) {
	require.False(t, results.Success().Failed())
	require.True(t, results.Failed("Context", "failed by").Failed())
}

func TestMessage(t *testing.T) {
	require.Equal(t, "", results.Success().Message())
	require.Equal(t, "Context: fail message", results.Failed("Context", "fail message").Message())
	require.Equal(t, "Context: fail message", results.Failed(" Context ", " fail message ").Message())
}
