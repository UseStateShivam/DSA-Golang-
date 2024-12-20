package main

func MaxArea(height []int) int {
	area := 0
	l, r := 0, len(height) - 1
	if len(height) < 2 {
		return 0
	}
	for l < r {
		area = max(min(height[l], height[r])*(r-l), area)
		if height[l] <= height[r] {
            l++
        } else {
            r--
        }
	}
	return area
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// 2ms
// 9.54mb