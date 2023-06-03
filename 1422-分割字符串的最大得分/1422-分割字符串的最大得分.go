package main

// leetcode 1422. 分割字符串的最大得分

func maxScore(s string) int {
    res := 0
    if len(s) == 0 {
        return 0
    }
    // 字符串s[1:]全部分配到右边
    left := 0
    if s[0] == '0' {
        left = 1
    }

    // 查看右边‘1’的个数，就是当前分配方案的最大结果
    right := 0
    for i:=1;i<len(s);i++{
        if s[i] == '1'{
            right++
        }
    }
    //fmt.Println("left: ",left)
    //fmt.Println("right: ",right)
    // 将字符串挨个分配到左边,至少要分成两组，所有从第一种方案开始遍历
    res = left + right
    for i:= 1;i<len(s) - 1;i++{
        if s[i] == '0'{
            left++
        }
        if s[i] == '1' {
            right--
        }
        // 将字符串挨个分配到左边
        temp := left + right
        res = max(temp,res)
    }
    return res
}

func max(i,j int) int{
    if i > j {
        return i
    }
    return j
}
