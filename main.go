package main

import "fmt"

func reverseSlice(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice before reversed:", arr)
	reverseSlice(arr)
	fmt.Println("Slice after reversed:", arr)
}
