package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// CountGoodNodes counts nodes that are ≥ every node on their root path.
// DFS carrying the running maximum: the comparison is against the path's
// max so far, which also updates as the walk descends.
func CountGoodNodes(root *TreeNode) int {
	const minInt = -int(^uint(0)>>1) - 1
	var walk func(n *TreeNode, pathMax int) int
	walk = func(n *TreeNode, pathMax int) int {
		if n == nil {
			return 0
		}
		count := 0
		if n.Val >= pathMax {
			count = 1
			pathMax = n.Val
		}
		return count + walk(n.Left, pathMax) + walk(n.Right, pathMax)
	}
	return walk(root, minInt)
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(3, tn(1, tn(3, nil, nil), nil), tn(4, tn(1, nil, nil), tn(5, nil, nil)))
	fmt.Println(CountGoodNodes(root))
}
