package main

func MaxProfitII(prices []int) int {
    // we can add profit even if the currProfit is < prevProfit

    left := 0
    profit := 0

    for right := 1; right < len(prices); right++ {
        if prices[right] < prices[left] {
            left = right
        } else if prices[right] - prices[left] > 0 {
            profit += prices[right] - prices[left]
            left = right
        }
    }

    return profit

    // accepted in one go?????????? wthhhhhhhhhhhhhhhhhhh????????
}

// 0ms
// 5.01mb