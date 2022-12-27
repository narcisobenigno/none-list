package es

import "github.com/narcisobenigno/none-list/pkg/results"

type Version struct {
	version uint64
}

func ParseVersion(version uint64) (Version, results.Result) {
	if version < 1 {
		return Version{}, results.Failed("Event", "version should be greater than or equal to 1")
	}
	return Version{version: version}, results.Success()
}

func MustParseVersion(version uint64) Version {
	parsedVersion, result := ParseVersion(version)
	if result.Failed() {
		panic(result.Message())
	}

	return parsedVersion
}
