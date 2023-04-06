package main

import (
	"strconv"
	"strings"
)

// leetcode 93-复原 IP 地址
func restoreIpAddresses(s string) []string {
	res := []string{}
	var dfs func(sub []string, startIndex int)
	dfs = func(sub []string, startIndex int) {
		// 选择合法的结果添加到结果集
		if len(sub) == 4 && startIndex == len(s) {
			res = append(res, strings.Join(sub, "."))
		}
		// ip 已经选择四段，但是s没有选完
		if len(sub) == 4 && startIndex < len(s) {
			return
		}
		// 每一个ip地址有四段
		for length := 1; length < 4; length++ {
			// 选择越界
			if startIndex+length > len(s) {
				return
			}
			// "0.011.255.245"、"192.168.1.312" 类似于这种不和方的结果
			if length != 1 && s[startIndex] == '0' {
				return
			}
			str := s[startIndex : startIndex+length]
			if n, _ := strconv.Atoi(str); n > 255 {
				return
			}
			sub = append(sub, str)
			dfs(sub, startIndex+length)
			sub = sub[:len(sub)-1]
		}
	}
	dfs([]string{}, 0)
	return res
}
