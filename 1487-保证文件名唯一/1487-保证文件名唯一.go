package main

import (
	"strconv"
)

/*
给你一个长度为 n 的字符串数组 names 。你将会在文件系统中创建 n 个文件夹：在第 i 分钟，新建名为 names[i] 的文件夹。

由于两个文件 不能 共享相同的文件名，因此如果新建文件夹使用的文件名已经被占用，系统会以 (k) 的形式为新文件夹的文件名添加后缀，其中 k 是能保证文件名唯一的 最小正整数 。

返回长度为 n 的字符串数组，其中 ans[i] 是创建第 i 个文件夹时系统分配给该文件夹的实际名称。
*/

// 1487-保证文件名唯一
func getFolderNames(names []string) []string {
	ans := make([]string, len(names))
	// 记住当前文件名重复到了第几个
	index := make(map[string]int)
	for p, name := range names {
		i := index[name]
		// 如果是第一个，没有重复的存在
		if i == 0 {
			index[name] = 1
			ans[p] = name
			continue
		}
		for index[name+"("+strconv.Itoa(i)+")"] > 0 {
			i++
		}
		t := name + "(" + strconv.Itoa(i) + ")"
		ans[p] = t
		index[name] = i + 1
		index[t] = 1
	}
	return ans
}
