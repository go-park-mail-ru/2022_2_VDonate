package utils

func Empty[T comparable](args ...T) bool {
	for _, arg := range args {
		if arg == *new(T) {
			return true
		}
	}
	return false
}
