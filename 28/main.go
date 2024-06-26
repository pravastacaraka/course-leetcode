package main

import (
	"fmt"
	"regexp"
)

func strStr(haystack string, needle string) int {
	re := regexp.MustCompile(needle)
	idx := re.FindAllStringIndex(haystack, -1)

	if len(idx) < 1 || len(idx[0]) < 1 {
		return -1
	}

	return idx[0][0]
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}
