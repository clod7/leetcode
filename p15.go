package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)

	for k := 0; k < len(nums); k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i, j := k+1, len(nums)-1
		for i < j {
			if nums[i]+nums[j]+nums[k] < 0 {
				i++
			} else if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			} else {
				j--
			}
		}
	}

	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
