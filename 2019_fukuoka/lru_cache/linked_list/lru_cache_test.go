package lrucache

import (
	"testing"
)

func TestSolution(t *testing.T) {
    var ans, expectAns int
    lruCache := Constructor()

	lruCache.Put(5, 3)
    lruCache.Put(1, 2)

    ans = lruCache.Get(5)
    expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
    
    lruCache.Evict()

    ans = lruCache.Get(1)
    expectAns = -1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
    
    ans = lruCache.Remove(5)
    expectAns = 3
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
}
