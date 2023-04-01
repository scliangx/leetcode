package main

// leetcode 132-分割回文串 II
func minCut(s string) int {
	if len(s) <= 1 {
		return 0
	}
	/*
		dp[i] 表示0到i的最小分割数, 求dp[len(s)-1]
		如果j到i是回文串,则dp[j] = min(dp[k],dp[i])
	*/
	dp := make([]int, len(s)+1)
	dp[0] = -1
	dp[1] = 1
	for i := 1; i <= len(s); i++ {
		// s[i] 最多需要分割i-1次
		dp[i] = i - 1
		for j := 0; j < i; j++ {
			if isPali(s, j, i-1) {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(s)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 双指针 判断字符串s是否回文
func isPali(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
