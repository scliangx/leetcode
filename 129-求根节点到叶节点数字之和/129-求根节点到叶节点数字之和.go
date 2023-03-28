package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	node *TreeNode
	num  int
}

// leetcode 129-求根节点到叶节点数字之和
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	q := []Pair{{root, root.Val}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		left := cur.node.Left
		right := cur.node.Right
		num := cur.num
		if left == nil && right == nil {
			sum += num
		}
		if left != nil {
			q = append(q, Pair{left, num*10 + left.Val})
		}
		if right != nil {
			q = append(q, Pair{right, num*10 + right.Val})
		}
	}
	return sum
}

// dfs实现
func sumNumbers1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(root, 0)
}

func helper(root *TreeNode, preSum int) int {
	if root == nil {
		return 0
	}
	sum := preSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return helper(root.Left, sum) + helper(root.Right, sum)
}
