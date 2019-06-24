package maxproductofthree

import (
	"sort"
)

// Time Complexity: O(nlogn)
// Space Complexity: O(1)
//
// Reference: http://bit.ly/2Y9dBI5
func Solution(A []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	// Max product of three numbers can be:
	//	1. Three largest numbers.
	//	2. Two smallest number (negative numbers) and one largest number.
	return max(A[0]*A[1]*A[2], A[len(A)-1]*A[len(A)-2]*A[0])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
