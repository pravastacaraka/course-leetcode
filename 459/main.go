package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	concatenated := s + s
	return strings.Contains(concatenated[1:len(concatenated)-1], s)
}

func main() {
	fmt.Println(repeatedSubstringPattern("babbabbabbabbab"))
}
