package utils

import "strings"

// TrimWhitespace will remove the tabs and newlines from a string
func TrimWhitespace(str string) string {
	r := strings.Replace(str, "\t", "", -1)
	r = strings.Replace(str, "\n", "", -1)
	return r
}
