package main

import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
	result := 0

	for _, char := range t {
		result += int(char)
	}

	for _, char := range s {
		result -= int(char)
	}

	return byte(result)
}

func main() {
	fmt.Println(findTheDifference("a", "aa"))
}
