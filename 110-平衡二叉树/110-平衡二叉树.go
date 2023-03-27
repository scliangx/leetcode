package main

// leetcode 110-平衡二叉树

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(depth(root.Left)-depth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	right := depth(root.Right)
	return 1 + max(left, right)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return 0 - x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
