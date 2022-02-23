package reversewords

import (
	"fmt"
	"strings"
)

func reverseusingstack(input string) (output string) {
	var stack Stack
	var result []string
	words := strings.Split(input, " ")
	fmt.Println("Words:", words)
	for _, word := range words {
		stack.Push(string(word))
	}
	for !stack.IsEmpty() {
		result = append(result, stack.Pop())
	}
	output = strings.Join(result, " ")
	fmt.Println("Reverse string:", output)
	return output
}

func reverseSimple(input string) (output string) {
	words := strings.Split(input, " ")
	fmt.Println("Words:", words)
	N := len(words)
	var result []string
	for i := N - 1; i >= 0; i-- {
		result = append(result, words[i])

	}
	output = strings.Join(result, " ")
	fmt.Println("Reverse string:", output)
	return output
}
func Reversefunc() {
	reverseusingstack("i love programming very much")
	reverseSimple("Hello how are you")
}
