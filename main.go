package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Diameter returns the number of edges on the longest path between any two
// nodes. The path through a node is leftHeight + rightHeight; the height
// recursion reports that for every node in the same single pass.
func Diameter(root *TreeNode) int {
	best := 0
	var height func(n *TreeNode) int
	height = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l, r := height(n.Left), height(n.Right)
		if l+r > best {
			best = l + r // longest path bending through n, in edges
		}
		if l > r {
			return l + 1
		}
		return r + 1
	}
	height(root)
	return best
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(1, tn(2, tn(4, nil, nil), tn(5, nil, nil)), tn(3, nil, nil))
	fmt.Println(Diameter(root))
}
