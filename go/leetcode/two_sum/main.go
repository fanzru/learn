package main

import (
	"fmt"
)

func main() {

	// two sum
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println("result test 1 : ", twoSum(nums1, target1))

	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println("result test 2 : ", twoSum(nums2, target2))

	nums3 := []int{3, 3}
	target3 := 6
	fmt.Println("result test 3 : ", twoSumMap(nums3, target3))

}

func twoSum(nums []int, target int) []int {
	lenNums := len(nums)
	for i := 0; i < lenNums; i++ {
		for j := i + 1; j < lenNums; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSumMap(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, found := numMap[complement]; found {
			return []int{idx, i}
		}
		numMap[num] = i
	}

	return []int{}
}
