package main

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func MinSubArrayLen(target int, nums []int) int {
	l, r, minLen := 0, 0, 0
    sum := nums[l]
	for r < len(nums) {
        if sum >= target {
            if minLen == 0 {
                minLen = r-l+1
            } else {
                minLen = min(minLen, r-l+1)
                if minLen == 1 {
                    return minLen
                }
            }
            sum -= nums[l]
            l++
        } else {
            if r != len(nums) - 1 {
                r++
                sum += nums[r]
            } else {
                return minLen
            }
        }
	}

	return minLen
}

// 0ms
// 9.42mb
// hahahah accepted in one go 🗿