package main

import "strings"

// leetcode 22-括号生成
func generateParenthesis(n int) []string {
	parenthesis := []string{}
	if n == 0 {
		return parenthesis
	}
	backtrack(n, 0, 0, []string{}, &parenthesis)
	return parenthesis
}

func backtrack(n, left, right int, s []string, parenthesis *[]string) {
	if len(s) == 2*n {
		*parenthesis = append(*parenthesis, strings.Join(s, ""))
		return
	}
	if left < n {
		s = append(s, "(")
		backtrack(n, left+1, right, s, parenthesis)
		s = s[:len(s)-1]
	}
	// `(` 一定多于 `)`，所以当前字符串中，right一定是小于left的
	if right < left {
		s = append(s, ")")
		backtrack(n, left, right+1, s, parenthesis)
		s = s[:len(s)-1]
	}
}
