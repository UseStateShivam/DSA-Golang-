package main

func Candy(ratings []int) int {
	c := make([]int, len(ratings))
    r := 0
	for i := range c {
		c[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			c[i] = c[i-1] + 1
		}
	}
	for i := len(ratings)-2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			c[i] = max(c[i+1]+1, c[i])
		}
	}
	for _, e := range c {
        r += e
    }
    return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 4ms
// 8.28mb