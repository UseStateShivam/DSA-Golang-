package main

func CanCompleteCircuit(gas []int, cost []int) int {
    totalGas, totalCost := 0, 0
    
    for i := 0; i < len(gas); i++ {
        totalGas += gas[i]
        totalCost += cost[i]
    }

    if totalGas < totalCost {
        return -1
    }

    tank, start := 0, 0

    for i := 0; i < len(gas); i++ {
        tank += gas[i] - cost[i]
        if tank < 0 {
            start = i + 1
            tank = 0
        }
    }

    return start
}

// 0ms
// 9.99ms
// BUT I AM FKING DONE