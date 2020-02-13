package phone

import (
	"log"
	"regexp"
	"strings"
	"unicode"
)

func Sanitize(s string) (string, bool) {
	if strings.HasPrefix(s, "+00") || strings.HasPrefix(s, "+ ") {
		return "", false
	}

	s = strings.TrimLeftFunc(s, func(r rune) bool {
		return string(r) == "0" || !unicode.IsNumber(r)
	})

	s = removeWhiteSpaceStringsBuilder(s)

	reg, err := regexp.Compile("^[0-9]+$")
	if err != nil {
		log.Fatal(err)
	}

	r := reg.Match([]byte(s))
	if !r {
		return "", false
	}

	size := len(s)
	if size == 3 || (size >= 7 && size <= 12) {
		return s, true
	}

	return "", false
}

func removeWhiteSpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
