package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: mycli <text>")
		os.Exit(1)
	}

	inputText := os.Args[1]
	outputText := alternateCase(inputText)
	fmt.Println(outputText)
}

func alternateCase(input string) string {
	var result strings.Builder
	upper := false

	for _, char := range input {
		if unicode.IsLetter(char) {
			if upper {
				result.WriteRune(unicode.ToUpper(char))
			} else {
				result.WriteRune(unicode.ToLower(char))
			}
			upper = !upper
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}
