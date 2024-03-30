package writing_cases

import (
	"strings"
	"unicode"
)

func ToWeirdCase(str string) string {
	words := strings.Fields(str)
	var resultWords []string

	for _, word := range words {
		var weirdWord []rune
		for i, v := range word {
			if i%2 == 0 {
				weirdWord = append(weirdWord, unicode.ToUpper(v))
			} else {
				weirdWord = append(weirdWord, unicode.ToLower(v))
			}
		}
		resultWords = append(resultWords, string(weirdWord))
	}

	return strings.Join(resultWords, " ")
}
