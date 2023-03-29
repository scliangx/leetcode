package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode 124-二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := math.MinInt32
	maxGain(root, &maxSum)
	return maxSum
}

func maxGain(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	// 递归计算左右子节点的最大贡献值
	// 只有在最大贡献值大于 0 时，才会选取对应子节点
	l := max(maxGain(root.Left, maxSum), 0)
	r := max(maxGain(root.Right, maxSum), 0)
	// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
	*maxSum = max(root.Val+l+r, *maxSum)
	// 返回节点的最大贡献值
	return root.Val + max(l, r)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
