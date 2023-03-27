package main

// leetcode 113-路径总和 II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	helper(root, targetSum, []int{}, &res)
	return res
}

func helper(root *TreeNode, targetSum int, paths []int, res *[][]int) {
	if root == nil {
		return
	}
	paths = append(paths, root.Val)
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		tmp := make([]int, len(paths))
		copy(tmp, paths)
		*res = append(*res, tmp)
	}
	helper(root.Left, targetSum-root.Val, paths, res)
	helper(root.Right, targetSum-root.Val, paths, res)
}
