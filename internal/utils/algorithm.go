package utils

import (
	"crypto/md5"
	"fmt"
)

func Empty[T comparable](args ...T) bool {
	var empty T
	for _, arg := range args {
		if arg == empty {
			return true
		}
	}
	return false
}

func GetMD5OfNumLast(data string, num int) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data[len(data)-num:])))
}
