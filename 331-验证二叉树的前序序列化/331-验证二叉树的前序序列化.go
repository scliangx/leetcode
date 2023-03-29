package main

// leetcode 331-验证二叉树的前序序列化
func isValidSerialization(preorder string) bool {
	if len(preorder) == 0 {
		return false
	}
	n := len(preorder)
	count := 1
	for i := 0; i < n; {
		if count == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			count--
			i++
		} else {
			for i < n && preorder[i] != ',' {
				i++
			}
			count++
		}
	}
	return count == 0
}
