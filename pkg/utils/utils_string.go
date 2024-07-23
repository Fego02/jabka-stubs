package utils

import "strings"

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(string(s[0])) + s[1:]
}

func IsFirstLetterLower(s string) bool {
	if len(s) == 0 {
		return false
	}
	firstLetter := string(s[0])
	firstLetter = strings.ToLower(firstLetter)

	return firstLetter == s
}
