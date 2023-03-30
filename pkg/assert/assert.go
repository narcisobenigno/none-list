package assert

func NoError(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](value T, err error) T {
	NoError(err)
	return value
}

func True(truthy bool, msg string) {
	if !truthy {
		panic(msg)
	}
}

func False(falsy bool, msg string) {
	True(!falsy, msg)
}
