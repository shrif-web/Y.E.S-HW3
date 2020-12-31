package internal

import "strings"

func IsEmpty(str string) bool {
	return len(strings.ReplaceAll(str, " ", "")) <= 0
}

func IsAllEmpty(strs ...string) bool {
	for _, str := range strs {
		if !IsEmpty(str) {
			return false
		}
	}
	return true
}
