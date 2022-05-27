package reverse_string

import "fmt"

func ReverseString(input string) (output string) {

	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
		fmt.Println(string(input[i]))
	}
	return output
}
