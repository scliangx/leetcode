package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode 257-二叉树的所有路径
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	constructPaths(root, "", &res)
	return res
}

func constructPaths(root *TreeNode, path string, paths *[]string) {
	if root != nil {
		t := path
		t += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			*paths = append(*paths, t)
		} else {
			t += "->"
			constructPaths(root.Left, t, paths)
			constructPaths(root.Right, t, paths)
		}
	}
}

type Paths struct {
	node *TreeNode
	path string
}

func binaryTreePaths1(root *TreeNode) []string {
	res := []string{}
	if root == nil {
		return res
	}
	q := []Paths{Paths{root, strconv.Itoa(root.Val)}}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			p := q[0]
			q = q[1:]
			if p.node.Left == nil && p.node.Right == nil {
				res = append(res, p.path)
			}
			if p.node.Left != nil {
				q = append(q, Paths{p.node.Left, p.path + "->" + strconv.Itoa(p.node.Left.Val)})
			}
			if p.node.Right != nil {
				q = append(q, Paths{p.node.Right, p.path + "->" + strconv.Itoa(p.node.Right.Val)})
			}
		}
	}
	return res
}
