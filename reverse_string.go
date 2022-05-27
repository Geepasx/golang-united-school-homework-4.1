package reverse_string

import (
	"strings"
)

func ReverseString(input string) (output string) {
	text := strings.Split(input, "")
	for i := len(text) - 1; i >= 0; i-- {
		output += string(text[i])
	}
	return output
}
