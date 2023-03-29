package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode 404-左叶子之和
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	dfs(root, &sum)
	return sum
}

func dfs(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*sum += root.Left.Val
	}
	dfs(root.Left, sum)
	dfs(root.Right, sum)
}
