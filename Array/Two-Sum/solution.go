package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		match := target - nums[i]

		if _, ok := m[match]; ok {
			return []int{m[match], i}
		}
		m[nums[i]] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 10, 11}
	target := 9

	ret := twoSum(nums, target)

	fmt.Println(ret)
}
