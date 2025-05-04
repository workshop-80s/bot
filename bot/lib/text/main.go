package text

import (
	"regexp"
	"strings"
)

func StripTag(s string) string {
	reg := regexp.MustCompile(`<.*?>`)
	return reg.ReplaceAllString(s, "${1}")
}

func Trim(s string) string {
	text := regexp.
		MustCompile(`\s+`).
		ReplaceAllString(s, " ")

	return strings.TrimSpace(text)
}
