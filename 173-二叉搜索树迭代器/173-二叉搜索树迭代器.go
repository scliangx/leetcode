package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	arr []int
}

// leetcode 173-二叉搜索树迭代器
func (it *BSTIterator) inorder(node *TreeNode) {
	if node == nil {
		return
	}
	it.inorder(node.Left)
	it.arr = append(it.arr, node.Val)
	it.inorder(node.Right)
}

func Constructor(root *TreeNode) (it BSTIterator) {
	it.inorder(root)
	return
}

func (this *BSTIterator) Next() int {
	val := this.arr[0]
	this.arr = this.arr[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.arr) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
