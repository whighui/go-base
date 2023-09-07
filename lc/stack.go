package lc

import "strings"

/**
栈先关操作呢
*/

// todo 华为面试  lc-402
// 给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。
// 示例 1 ：
// 输入：num = "1432219", k = 3
// 输出："1219"
// 解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219
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

//输入: temperatures = [73,74,75,71,69,72,76,73] 输出: [1,1,4,2,1,1,0,0]

// O(N^2)  for循环变量两次指定是能完成的 但是题目估计要求O(n)来完成呗
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(temperatures))
	for index := len(temperatures) - 1; index >= 0; index-- {
		for len(stack) > 0 && temperatures[index] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
		result[index] = stack[len(stack)-1] - index
	}
	return result
}
