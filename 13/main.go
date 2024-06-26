package main

import "fmt"

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var romanMapExclusive = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {
	var result int
	for i := len(s) - 1; i >= 0; i-- {
		var (
			prevChar = ""
			char     = string(s[i])
		)
		if i > 0 {
			prevChar = string(s[i-1])
		}

		if v, ok := romanMapExclusive[prevChar+char]; ok {
			result += v
			i--
		} else {
			result += romanMap[char]
		}
	}
	return result
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
