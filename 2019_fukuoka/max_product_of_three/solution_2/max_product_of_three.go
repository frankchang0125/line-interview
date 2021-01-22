package maxproductofthree

import (
	"math"
)

// Time Complexity: O(n)
// Space Complexity: O(1)
//
// Reference: http://bit.ly/2Y9dBI5
func Solution(A []int) int {
	max1 := math.MinInt32
	max2 := math.MinInt32
	max3 := math.MinInt32
	min1 := math.MaxInt32
	min2 := math.MaxInt32

	for _, num := range A {
		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num > max2 {
			max3 = max2
			max2 = num
		} else if num > max3 {
			max3 = num
		}

		if num < min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
	}

	return max(max1*max2*max3, min1*min2*max1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
