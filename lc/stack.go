package lc

import "strings"

/**
栈先关操作呢
*/

//todo 华为面试
//给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。
//示例 1 ：
//输入：num = "1432219", k = 3
//输出："1219"
//解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219
func removeKdigits(num string, k int) string {
	stack := make([]uint8, 0, len(num))

	for index := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[index] { //让数字以递增序列增长
			stack = stack[0 : len(stack)-1]
			k--
		}
		stack = append(stack, num[index])
	}

	//如若k还存在呢
	if k > 0 {
		stack = stack[0 : len(stack)-k]
	}

	result := strings.TrimLeft(string(stack), "0")

	if result == "" {
		return "0"
	}
	return result
}
