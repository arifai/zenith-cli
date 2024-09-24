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

func ConvertCase(str string, capitalizeFirst bool) string {
	var result strings.Builder
	upperNext := capitalizeFirst

	for _, r := range str {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			if r == '_' || r == ' ' {
				upperNext = true
			}
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
		} else if r == '_' {
			result.WriteRune('_')
		} else if unicode.IsUpper(r) {
			if i > 0 && (unicode.IsLower(rune(str[i-1])) || str[i-1] == '_') {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
