package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[j]+nums[j-i] == target {
				return []int{j - i, j}
			}
		}
	}
	return nil
}

func main() {
	var (
		nums   = []int{3, 2, 4}
		target = 6
	)

	fmt.Printf("result: %v\n", TwoSum(nums, target))
}
