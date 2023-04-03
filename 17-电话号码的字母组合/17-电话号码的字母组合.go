package main

// leetcode 17-电话号码的字母组合
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	combinations := []string{}
	if len(digits) == 0 {
		return combinations
	}
	backtrack(digits, 0, "", &combinations)
	return combinations
}

func backtrack(digits string, index int, combination string, phoneSlice *[]string) {
	if len(combination) == len(digits) {
		*phoneSlice = append(*phoneSlice, combination)
	} else {
		// 解析出来输入的数字
		digit := string(digits[index])
		// 获取输入数字的对应字符串
		letters := phoneMap[digit]
		// 当前对应字符串的长度
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]), phoneSlice)
		}
	}
}
