package main

// leetcode 131-分割回文串
func partition(s string) [][]string {
	res := [][]string{}
	if len(s) == 0 {
		return res
	}
	backtrack(s, []string{}, 0, &res)
	return res
}

func backtrack(s string, tmp []string, index int, res *[][]string) {
	if index == len(s) {
		t := make([]string, len(tmp))
		copy(t, tmp)
		*res = append(*res, t)
		return
	}
	for i := index; i < len(s); i++ {
		if isPali(s, index, i) {
			tmp = append(tmp, s[index:i+1])
			// backtrack
			backtrack(s, tmp, i+1, res)
			// 剪枝
			tmp = tmp[:len(tmp)-1]
		}
	}
	return
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
