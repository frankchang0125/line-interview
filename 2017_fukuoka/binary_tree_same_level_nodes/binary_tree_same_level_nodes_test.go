package binary_tree_same_level_nodes

import (
	"reflect"
	"testing"
)

func TestLeftNodesOnly(t *testing.T) {
	//          1
	//         /
	//        2
	//       /
	//      3
	//     /
	//    5
	tree := &Tree{
		Root: &Node{
			Val: 1,
			Left: &Node{
				Val: 2,
				Left: &Node{
					Val: 3,
					Left: &Node{
						Val: 4,
						Left: &Node{
							Val: 5,
						},
					},
				},
			},
		},
	}

	ans := [][]int{
		[]int{1},
		[]int{2},
		[]int{3},
		[]int{4},
		[]int{5},
	}

	result := tree.sameLevelNodes()
	if !reflect.DeepEqual(result, ans) {
		t.Errorf("expect ans to be %v, but got %v", ans, result)
	}
}

func TestRightNodesOnly(t *testing.T) {
	//          1
	//           \
	//            2
	//             \
	//              3
	//               \
	//                4
	//                 \
	//                  5
	tree := &Tree{
		Root: &Node{
			Val: 1,
			Right: &Node{
				Val: 2,
				Right: &Node{
					Val: 3,
					Right: &Node{
						Val: 4,
						Right: &Node{
							Val: 5,
						},
					},
				},
			},
		},
	}

	ans := [][]int{
		[]int{1},
		[]int{2},
		[]int{3},
		[]int{4},
		[]int{5},
	}

	result := tree.sameLevelNodes()
	if !reflect.DeepEqual(result, ans) {
		t.Errorf("expect ans to be %v, but got %v", ans, result)
	}
}

func TestFullNodes(t *testing.T) {
    //                 1
    //           /           \
    //          2             3
    //       /     \       /     \
    //      4       5     6       7
    //     / \     / \   / \     / \
    //    8   9   10 11 12 13   14 15
	tree := &Tree{
		Root: &Node{
			Val: 1,
			Left: &Node{
				Val: 2,
				Left: &Node{
					Val: 4,
					Left: &Node{
						Val: 8,
                    },
                    Right: &Node{
                        Val: 9,
                    },
				},
				Right: &Node{
					Val: 5,
					Left: &Node{
						Val: 10,
					},
					Right: &Node{
						Val: 11,
					},
				},
            },
            Right: &Node{
                Val: 3,
                Left: &Node{
					Val: 6,
					Left: &Node{
						Val: 12,
                    },
                    Right: &Node{
                        Val: 13,
                    },
				},
				Right: &Node{
					Val: 7,
					Left: &Node{
						Val: 14,
					},
					Right: &Node{
						Val: 15,
					},
				},
            },
		},
	}

	ans := [][]int{
		[]int{1},
		[]int{2, 3},
		[]int{4, 5, 6, 7},
		[]int{8, 9, 10, 11, 12, 13, 14, 15},
	}

	result := tree.sameLevelNodes()
	if !reflect.DeepEqual(result, ans) {
		t.Errorf("expect ans to be %v, but got %v", ans, result)
	}
}

func TestNodes(t *testing.T) {
	//          1
	//         /
	//        2
	//       / \
	//      3   4
	//     /   /  \
	//    5   6    7
	tree := &Tree{
		Root: &Node{
			Val: 1,
			Left: &Node{
				Val: 2,
				Left: &Node{
					Val: 3,
					Left: &Node{
						Val: 5,
					},
				},
				Right: &Node{
					Val: 4,
					Left: &Node{
						Val: 6,
					},
					Right: &Node{
						Val: 7,
					},
				},
			},
		},
	}

	ans := [][]int{
		[]int{1},
		[]int{2},
		[]int{3, 4},
		[]int{5, 6, 7},
	}

	result := tree.sameLevelNodes()
	if !reflect.DeepEqual(result, ans) {
		t.Errorf("expect ans to be %v, but got %v", ans, result)
	}
}
