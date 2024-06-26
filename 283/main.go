package main

func moveZeroes(nums []int) {
	n := 0
	for _, v := range nums {
		if v != 0 {
			nums[n] = v
			n++
		}
	}
	for i := n; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
