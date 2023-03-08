package main

/*
给你一个大小为 n x n 的整数矩阵 grid 。
生成一个大小为 (n - 2) x (n - 2) 的整数矩阵  maxLocal ，并满足：
maxLocal[i][j] 等于 grid 中以 i + 1 行和 j + 1 列为中心的 3 x 3 矩阵中的 最大值 。
换句话说，我们希望找出 grid 中每个 3 x 3 矩阵中的最大值。
返回生成的矩阵
*/

// 暴力破解
// 2373-矩阵中的局部最大值
func largestLocal(grid [][]int) [][]int {
	if len(grid) == 0 || grid == nil {
		return [][]int{}
	}
	n := len(grid)
	maxLocal := make([][]int, n-2)
	for i := 1; i < n-1; i++ {
		rows := make([]int, n-2)
		for j := 1; j < n-1; j++ {
			tmpMax := grid[i][j]
			// 遍历以当前位置为中心，3*3范围内最大的值
			for c := i - 1; c <= i+1; c++ {
				for c2 := j - 1; c2 <= j+1; c2++ {
					if tmpMax < grid[c][c2] {
						tmpMax = grid[c][c2]
					}
				}
			}
			rows[j-1] = tmpMax
		}
		maxLocal[i-1] = rows
	}
	return maxLocal
}
