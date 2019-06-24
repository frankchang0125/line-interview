package slidingwindowmaximum

import (
    "reflect"
	"testing"
)

func TestSolution(t *testing.T) {
    var ans, expectAns []int

	ans = maxSlidingWindow([]int{2, 1, 2, -1, 3}, 2)
    expectAns = []int{2, 2, 2, 3}
    if !reflect.DeepEqual(ans, expectAns) {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
