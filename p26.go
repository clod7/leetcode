package main

import "log"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 1
	temp := nums[0]
	for _, num := range nums {
		if num != temp {
			nums[count] = num
			temp = num
			count ++
		}
	}

	nums = nums[:count]
	return len(nums)
}

func main() {
	log.Println(removeDuplicates([]int{1, 1, 2}))
}
