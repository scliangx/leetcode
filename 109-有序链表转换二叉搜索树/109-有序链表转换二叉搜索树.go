package main

// leetcode 109-有序链表转换二叉搜索树

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	return helper(head, nil)
}

func getMiddNode(left, right *ListNode) *ListNode {
	slow, fast := left, left
	for fast != right && fast.Next != right {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func helper(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	midNode := getMiddNode(left, right)
	root := &TreeNode{Val: midNode.Val}
	root.Left = helper(left, midNode)
	root.Right = helper(midNode.Next, right)
	return root
}
