//给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现
//在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//
//
// 示例 1:
//
//
//输入: temperatures = [73,74,75,71,69,72,76,73]
//输出: [1,1,4,2,1,1,0,0]
//
//
// 示例 2:
//
//
//输入: temperatures = [30,40,50,60]
//输出: [1,1,1,0]
//
//
// 示例 3:
//
//
//输入: temperatures = [30,60,90]
//输出: [1,1,0]
//
//
//
// 提示：
//
//
// 1 <= temperatures.length <= 10⁵
// 30 <= temperatures[i] <= 100
//
//
// Related Topics 栈 数组 单调栈 👍 1599 👎 0

package doc

//这里边就是判断高温天气呗 如何做呢

// leetcode submit region begin(Prohibit modification and deletion)
// O(N^2)  for循环变量两次指定是能完成的 但是题目估计要求O(n)来完成呗
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(temperatures))
	//倒序 逆向思维
	for index := len(temperatures) - 1; index >= 0; index-- {
		for len(stack) > 0 && temperatures[index] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			result[index] = stack[len(stack)-1] - index //存放索引 不存放值
		} else {
			result[index] = 0
		}
		stack = append(stack, index)  //最后才是放索引

	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
