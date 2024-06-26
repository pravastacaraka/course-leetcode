package main

import (
	"fmt"
	"unicode"
)

func toLowerCase(s string) string {
	var result string
	for _, c := range s {
		if unicode.IsLetter(c) {
			result += string(c + 32)
		}
	}
	return result
}

func main() {
	fmt.Println(toLowerCase("LOVELY"))
}
