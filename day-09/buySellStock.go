package main

func MaxProfit(prices []int) int {
	// given: prices array
	// return: profit

	// approach
	// 2 pointer
	// keep left pointer on 0th index
	// right pointer on last index
	// check for every step if element on right index > left index
	// if yes then update the profit variable if currProfit > prevProfit
	// move the right pointer towards left until left pointer is met
	// then reset the right pointer to last index
	// and left index to +1
	// check until left meets right
	// return profit

	// left := 0
	// right := len(prices) - 1
	// profit := 0

	// for right != left {
	//     if prices[right] - prices[left] > profit {
	//         profit = prices[right] - prices[left]
	//     }
	//     right--
	//     if right == left {
	//         right = len(prices) - 1
	//         left++
	//     }
	// }

	// return profit

	// works but uses O(n^2) complexity

	left := 0
	profit := 0

	for right := 1; right < len(prices); right++ {
		if prices[right] < prices[left] {
			left = right
		} else if prices[right] - prices[left] > profit {
            profit = prices[right] - prices[left]
		}
	}

	return profit
}

// 0ms
// 13mb :(