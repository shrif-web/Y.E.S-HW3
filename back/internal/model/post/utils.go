package post

import "strings"

func isEmpty(str string) bool {
	return len(strings.ReplaceAll(str, " ", "")) <= 0
}

func isAllEmpty(strs ...string) bool {
	for _, str := range strs {
		if !isEmpty(str) {
			return false
		}
	}
	return true
}
