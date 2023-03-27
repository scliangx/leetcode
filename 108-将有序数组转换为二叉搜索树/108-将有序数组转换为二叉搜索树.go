package main

// leetcode 108-将有序数组转换为二叉搜索树

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 || nums == nil {
		return nil
	}
	return generate(nums, 0, len(nums)-1)
}

func generate(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end + 1) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = generate(nums, start, mid-1)
	root.Right = generate(nums, mid+1, end)
	return root
}
