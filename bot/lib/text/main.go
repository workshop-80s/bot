package text

import "regexp"

func StripTag (s string) string {
	reg := regexp.MustCompile(`<.*?>`)
	return reg.ReplaceAllString(s, "${1}")
}