package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	temp := word1 + word2
	var merged string
	for i := range temp {
		if i > len(word1)-1 {
			merged += string(word2[i:])
			break
		} else if i > len(word2)-1 {
			merged += string(word1[i:])
			break
		} else {
			merged += string(temp[i]) + string(temp[len(word1)+i])
		}
	}
	return merged
}

func main() {
	fmt.Println(mergeAlternately("abcd", "pq"))
}
