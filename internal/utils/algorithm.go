package utils

func Empty[T comparable](args ...T) bool {
	var empty T
	for _, arg := range args {
		if arg == empty {
			return true
		}
	}
	return false
}
