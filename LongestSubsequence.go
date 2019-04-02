package main

import (
	"fmt"
	"runtime"
)

// Get the max integer value
func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func main() {
	runtime.GOMAXPROCS(4)

	nums := []int{-3, 4, -3, 1, 2, 3, 10, -9, 7}

	var maxValue = nums[0]
	dp := make([]int, len(nums))

	var start, end int

	// Calculate dp[1..n]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if maxValue < dp[i] {
			maxValue = dp[i]
			end = i
		}
	}

	// Find Subsequence's start and end position
	var copyMaxValue = maxValue
	for i := end; i >= 0; i-- {
		copyMaxValue -= nums[i]
		if copyMaxValue == 0 {
			start = i
			break
		}
	}

	fmt.Println(maxValue, " start=", start, " end=", end)
}
