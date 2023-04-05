package main

// leetcode 79-单词搜索
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word[0] == board[i][j] && backtrack(board, word, used, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, used [][]bool, i, j, k int) bool {
	// 每个字符均匹配，则找到了答案
	if k == len(word) {
		return true
	}
	// 越界
	if (i < 0 || i >= len(board)) || (j < 0 || j >= len(board[0])) {
		return false
	}
	// 不允许重复访问
	if used[i][j] {
		return false
	}
	// 不等，直接返回
	if board[i][j] != word[k] {
		return false
	}
	used[i][j] = true
	isExist := backtrack(board, word, used, i+1, j, k+1) ||
		backtrack(board, word, used, i-1, j, k+1) ||
		backtrack(board, word, used, i, j+1, k+1) ||
		backtrack(board, word, used, i, j-1, k+1)

	used[i][j] = false
	return isExist
}
