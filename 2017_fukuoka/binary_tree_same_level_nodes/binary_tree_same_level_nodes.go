package binary_tree_same_level_nodes

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) sameLevelNodes() [][]int {
	return bfs(t.Root)
}

func bfs(root *Node) [][]int {
	ans := [][]int{}

	if root == nil {
		return ans
	}

	queue := []*Node{root} // Enqueue

	for len(queue) > 0 {
		values := []int{}

		for i := len(queue); i > 0; i-- {
			n := queue[0]
			queue = queue[1:] // Dequeue

			values = append(values, n.Val)

			if n.Left != nil {
				queue = append(queue, n.Left) // Enqueue
			}

			if n.Right != nil {
				queue = append(queue, n.Right) // Enqueue
			}
		}

		ans = append(ans, values)
	}

	return ans
}
