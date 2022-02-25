package must

// Must returns val if err is nil, otherwise it panics.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
