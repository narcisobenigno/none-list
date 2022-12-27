package es

import "github.com/narcisobenigno/none-list/pkg/results"

type Version struct {
	version uint
}

func ParseVersion(version uint) (Version, results.Result) {
	if version < 1 {
		return Version{}, results.Failed("Event", "version should be greater than or equal to 1")
	}
	return Version{version: version}, results.Success()
}

func MustParseVersion(version uint) Version {
	parsedVersion, result := ParseVersion(version)
	if result.Failed() {
		panic(result.Message())
	}

	return parsedVersion
}

func InitialVersion() Version {
	return MustParseVersion(1)
}
