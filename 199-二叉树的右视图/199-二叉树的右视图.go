package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode 199-二叉树的右视图
// bfs
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		item := 0
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			item = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, item)
	}
	return res
}

// dfs
func rightSideView1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	DFS(root, 0, &res)
	return res
}

func DFS(root *TreeNode, depth int, res *[]int) {
	if root == nil {
		return
	}
	// 右子树，左子树 顺序访问，第一个被访问到的就是右视图
	if len(*res) == depth {
		*res = append(*res, root.Val)
	}
	// 树的深度判断是不是第一个被访问
	depth++
	DFS(root.Right, depth, res)
	DFS(root.Left, depth, res)
}
