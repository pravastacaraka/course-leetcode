package main

import "fmt"

func signFunc(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func arraySign(nums []int) int {
	var temp int
	for _, v := range nums {
		if v == 0 {
			return signFunc(0)
		}
		if temp == 0 {
			temp = signFunc(v)
		} else {
			temp *= signFunc(v)
		}
	}
	return temp
}

func main() {
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
}
