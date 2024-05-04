package utils

import "strings"

func IsEmptyString(str string) bool {
	if len(strings.TrimSpace(str)) == 0 {
		return true
	}
	return false
}
