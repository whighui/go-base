package lc

import (
	"math"
)

/**
动态规划
1.确定dp数组（dp table）以及下标的含义
2.确定递推公式
3.dp数组如何初始化
4.确定遍历顺序
5.举例推导dp数组
*/

//01背包 一个物品只能拿一次 给定背包容量  计算最大值
//物品  重量 价值
//物品0	1	15
//物品1	3	20
//物品2	4	30

//01背包
//
func packet01(weight []int, value []int, bagSize int) int {
	//dp[i][j] 指下标【0--i】之间任意取，放进容量为j的背包价值最大
	dp := make([][]int, len(weight)) //i代表物品的数量  j 代表背包容量
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, bagSize+1) //每一行初始化背包容量
	}
	//背包容量初始化
	for j := weight[0]; j <= bagSize; j++ {
		dp[0][j] = weight[0]
	}

	//遍历二维数组
	for i := 1; i < len(weight); i++ { //遍历物品
		for j := 1; j <= bagSize; j++ { //遍历背包
			if j >= weight[i-1] { //因为dp数组的定义 i-1对应物品i
				//可以选择放入物品和不选择放入物品
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i]) //放入物品和不放入物品两种取最大值呗
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(weight)-1][bagSize]
}

/**
--------------------------------------------------------------------------------------------------------------------完全背包
组合遍历物品
排列遍历背包
*/

//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//属于完全背包范围 ----> 并且属于组合问题 所以遍历先遍历物品还是先遍历背包都是OK的
func coinChangeVersion(coins []int, amount int) int {
	dp := make([]int, amount+1) //以金额i为结尾 组成硬币所需最小数量
	for index := range dp {
		dp[index] = math.MaxInt //因为在遍历过程当中需要找到最小值，所以需要进行初始化成最大值
	}

	dp[0] = 0 //凑成金额0 所需硬币数量0

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}

		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]

}

//零钱兑换II 凑成硬币金额的组合数
//输入: amount = 5, coins = [1, 2, 5]
//输出: 4

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//完全背包 一个物品可以拿多次 计算最大值
func packetTotally() {
	//一维数组正序遍历 可以重复添加
	//01背包 一维数组倒序遍历 防止重复添加
}

//完全背包

//买卖股票 I II III IIII

//小偷偷零钱 I II III

//俄罗斯套娃信封问题

//数组最长递增子序列

//数组最长递增子序列的个数

/**
//lc-53 最大子数组和  子数组和代表是子数组是连续的
//todo 1.使用滑动窗口  2.使用动态规划算法  题目子数组最大和就是一定要返回一个值呢
*/
func maxSubArrayWithWindown(nums []int) int {
	//滑动窗口 关键点就是何时来放大窗口 何时来缩小窗口 进行左右指针进行移动
	left, right := 0, 0
	maxSum, window := math.MinInt, 0 //maxSum代表滑动窗口最大值
	/**
	这里边定义一个窗口：
	*/
	for left <= right && right < len(nums) {
		window = window + nums[right]
		right++
		maxSum = max(window, maxSum)      //这里边最大子数组和一定就是要返回一个值
		for left <= right && window < 0 { //在这里应该是小于0 才能进行缩短窗口
			window = window - nums[left]
			left++
		}
	}

	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	result := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i]) //子数组连续  子数组不连续以当前nums[i]就是当前子数组
		//[-1,-2,-3,4,-5]
		result = max(result, dp[i])
	}

	return result
}

//数组最长递增子序列  子序列在这里就是不连续的
//dp[i]定义===以nums[i]为结尾的数组最长递增子序列  最后在遍历dp[i]找到最大值
//dp[i]=dp[i-1] 怎么来进行推导呢 因为这里边就是不连续的 在这里就是如何进行判断呢
//dp[i-1]如何推导dp[i]
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	dp := make([]int, len(nums)) //确定dp是以nums[i]为结尾的最长递增子序列
	for i := 0; i < len(nums); i++ {
		dp[i] = 1 //每个元素都可以当做递增子序列
	}
	result := dp[0]

	for i := 1; i < len(nums); i++ {
		//以dp[i]为结尾的最长递增子序列  以nums[i]为结尾的最长递增子序列 结尾 重点是结尾
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1) //dp初始化都是1  因为每个元素都可以当做当前递增子序列呢
				result = max(result, dp[i])
			}
		}
	}
	return result
}

//todo  数组最长递增子序列的个数  拼多多二面面试
//最长递增子序列的个数  【1，1，1，1，1】 个数为5个 最长递增子序列为1 长度是1 总体个数为5个
//[1,2,3,5,4,7]  返回2
//存在两个最长递增子序列 1，2，3，5，7   1，2，3，4，7
func lengthOfLISCount(nums []int) int {
	//dp[i]  以nums[i]为结尾的最长递增子序列的长度
	//cnt[i] 以nums[i]为结尾的最长递增子序列个数  但是这里边个数就是需要进行计数 来尝试做一做

	if len(nums) <= 0 {
		return len(nums)
	}
	dp, cnt := make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		cnt[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				} else if dp[j]+1 == dp[i] {
					cnt[i] = cnt[j] + cnt[j]
				}
			}
		}
	}
	//找到dp[i]最大值的索引 在找到

	maxLength := dp[0] //这里边maxLength可以在db数组循环遍历中获取
	for i := 1; i < len(nums); i++ {
		if maxLength < dp[i] {
			maxLength = dp[i]
		}
	}

	result := 0
	//如果dp[i]=[1,1,1,1] 那么最长递增子序列长度为1  个数为4个 所以还要遍历个数数组
	for i := 1; i < len(nums); i++ {
		if maxLength == dp[i] {
			result = result + cnt[i]
		}
	}

	return result
}

//俄罗斯套娃信封问题
