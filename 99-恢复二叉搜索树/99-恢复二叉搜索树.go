package main

// leetcode 99-恢复二叉搜索树

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	res := []*TreeNode{}
	treeHelper(root, &res)
	x, y := -1, -1
	for i := 0; i < len(res)-1; i++ {
		if res[i].Val > res[i+1].Val {
			// 后一个找找到小的，所以是i+1
			y = i + 1
			if x == -1 {
				// 前一个找到大的，所以是i
				x = i
			}
		}
	}
	res[x].Val, res[y].Val = res[y].Val, res[x].Val
	return
}

func treeHelper(root *TreeNode, res *[]*TreeNode) {
	if root == nil {
		return
	}
	treeHelper(root.Left, res)
	*res = append(*res, root)
	treeHelper(root.Right, res)
}
