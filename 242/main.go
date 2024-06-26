package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var temp = make(map[rune]int)
	for _, val := range s {
		temp[val]++
	}

	for _, val := range t {
		temp[val]--
	}

	for _, v := range temp {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("ac", "bb"))
}
