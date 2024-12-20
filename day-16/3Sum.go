package main

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r]+nums[i] == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if nums[l]+nums[r]+nums[i] < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// 18ms
// 9.65mb