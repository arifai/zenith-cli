package utils

import (
	"os"
	"strings"
	"unicode"
)

func CheckGoModFileExists() bool {
	_, err := os.Stat("go.mod")
	return !os.IsNotExist(err)
}

func ToCamelCase(str string) string {
	var result strings.Builder
	upperNext := true
	for _, r := range str {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			continue
		}
		if r == '_' || r == ' ' {
			upperNext = true
			continue
		}
		if upperNext {
			result.WriteRune(unicode.ToUpper(r))
			upperNext = false
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}
	return result.String()
}

func ToSnakeCase(str string) string {
	var result strings.Builder
	for i, r := range str {
		if r == ' ' {
			result.WriteRune('_')
		} else if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			continue
		} else if unicode.IsUpper(r) {
			if i > 0 && result.Len() > 0 && result.String()[result.Len()-1] != '_' {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
