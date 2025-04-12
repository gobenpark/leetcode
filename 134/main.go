package main

import "fmt"

// TODO: 복습
// https://leetcode.com/problems/gas-station/?envType=problem-list-v2&envId=greedy
func canCompleteCircuit(gas []int, cost []int) int {

	totalGas := 0
	totalCost := 0
	startPoint := 0
	tank := 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]

		tank += gas[i] - cost[i]
		if tank < 0 {
			startPoint = i + 1
			tank = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}

	return startPoint
}

func main() {

	result := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	fmt.Println(result)
}
