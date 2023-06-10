package main

// leetcode 1219-黄金矿工

func getMaximumGold(grid [][]int) int {
    res := 0
    if len(grid) == 0 || grid == nil {
        return res
    }
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[i]);j++{
            if grid[i][j] == 0 {
                continue
            }
            backtrack(grid,i,j,0,&res)
        }
    }
    return res
}


func backtrack(grid [][]int,row int,col int,cur int, res *int) {
    if row < 0 || row >= len(grid) {
        return
    }
    if col < 0 || col >= len(grid[row]) {
        return
    }

    if grid[row][col] == 0 {
        return
    }
    cur += grid[row][col]
    if cur > *res {
        *res = cur
    }
    
    temp := grid[row][col]
    grid[row][col] = 0
    backtrack(grid,row+1,col,cur,res)
    backtrack(grid,row-1,col,cur,res)
    backtrack(grid,row,col+1,cur,res)
    backtrack(grid,row,col-1,cur,res)
    grid[row][col] = temp

}
