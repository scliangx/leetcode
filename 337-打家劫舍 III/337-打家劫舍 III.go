package main

// leetcode 337-打家劫舍 III

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
f(o) 表示选择 o 节点的情况下，o 节点的子树上被选择的节点的最大权值和；
g(o) 表示不选择 o 节点的情况下，o 节点的子树上被选择的节点的最大权值和；
l 和 r 代表 o 的左右孩子
*/
func rob(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

// 返回值长度为2，idnex=0表示选择，index=1表示不选择
func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
