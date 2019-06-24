package mazeshortestpath

import (
	"testing"
)

func TestSolution(t *testing.T) {
    var ans, expectAns int
    var maze [][]byte

    maze = [][]byte{
        []byte("11111"),
        []byte("1S0G1"),
        []byte("11111"),
    }
	ans = goMaze(maze, 5, 3)
	expectAns = 2
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	maze = [][]byte{
        []byte("11111111"),
        []byte("1S000001"),
        []byte("11101101"),
        []byte("1G000001"),
        []byte("11111111"),
    }
	ans = goMaze(maze, 8, 5)
	expectAns = 6
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
