package main

/**
 * Definition for a Node.
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 116-填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		// tmp实际是当前层的数据
		tmp := queue
		queue = nil
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

// 递归思路
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	// 连接左子树
	connectTwoNode(node1.Left, node1.Right)
	// 连接右子树
	connectTwoNode(node2.Left, node2.Right)
	// 连接不同父结点
	connectTwoNode(node1.Right, node2.Left)
}

// 遍历
func connect3(root *Node) *Node {
	if root == nil {
		return root
	}
	// 每次循环从该层的最左侧节点开始
	for cur := root; cur.Left != nil; cur = cur.Left {
		// 通过 Next 遍历这一层节点，为下一层的节点更新 Next 指针
		for node := cur; node != nil; node = node.Next {
			// 左节点指向右节点
			node.Left.Next = node.Right
			// 右节点指向下一个左节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}
