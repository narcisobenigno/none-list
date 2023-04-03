package es

import (
	"github.com/narcisobenigno/grocery-go/pkg/assert"
	"github.com/narcisobenigno/grocery-go/pkg/results"
)

type Version struct {
	version uint
}

func TryParseVersion(version uint) (Version, results.Result) {
	if version < 1 {
		return Version{}, results.Failed("Event", "version should be greater than or equal to 1")
	}
	return Version{version: version}, results.Success()
}

func ParseVersion(version uint) Version {
	parsedVersion, result := TryParseVersion(version)

	assert.False(result.Failed(), result.Message())

	return parsedVersion
}

func InitialVersion() Version {
	return ParseVersion(1)
}
