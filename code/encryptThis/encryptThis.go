package encryptThis

import (
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	var result = ""
	if text == "" {
		return ""
	}
	var words = strings.Split(text, " ")
	for wordKey, word := range words {
		result += strconv.Itoa(int(word[0]))
		if len(word) > 1 {
			result += string(word[len(word)-1])
		}
		if len(word) == 3 {
			result += string(word[1])
		}
		if len(word) > 3 {
			result += word[2:len(word)-1] + string(word[1])
		}
		if wordKey != len(words)-1 {
			result += " "
		}
	}
	return result
}
