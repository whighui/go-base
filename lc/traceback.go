package lc

import (
	"sort"
	"strconv"
)

//---------------------------------------排列+组合

// lc-46  全排列
func permute(nums []int) [][]int {
	if len(nums) == 0 || nums == nil {
		return nil
	}
	result, list, visited := make([][]int, 0), make([]int, 0), make([]bool, len(nums))

	var trace func(index int)
	trace = func(index int) {
		if len(list) == len(nums) {
			temp := make([]int, len(list))
			copy(temp, list)
			result = append(result, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			list = append(list, nums[i])
			trace(i + 1)
			visited[i] = false
			list = list[:len(list)-1] //这里边就是去除上一个元素呗 因为这里边就是全排列 就是跟全组合并不是一样的呗
		}
	}
	trace(0)
	return result
}

// 全排列II 带有重复数字 所以这里边就是需要进行减枝被
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 || nums == nil {
		return nil
	}

	result, list, visited := make([][]int, 0), make([]int, 0), make([]bool, len(nums))
	var traceback func(index int)
	sort.Ints(nums) //返回nums是以递增序列来返回呢
	traceback = func(index int) {
		if index == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, list)
			result = append(result, list)
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] || i > 0 && nums[i] == nums[i-1] {
				continue
			}

			list = append(list, nums[i])
			visited[i] = true
			traceback(i + 1)
			visited[i] = false
			list = list[:len(list)-1]
		}
	}

	traceback(0)
	return result
}

// lc-60 第K个排列
// leetcode submit region begin(Prohibit modification and deletion)
func getPermutation(n int, k int) string {
	/**
	首先如果按照正常全排列 就是找到第K个位置进行排列  时间复杂度也是比较高
	但是如果在这里就是按照数学思想固定一个范围 找到第K个位置不是更好吗 就是需要找到起始位置
	*/

	group := 1
	for i := 1; i <= n; i++ {
		group = group * i
	}
	result, visited := make([]string, 0), make([]bool, n)
	var dfs func()
	dfs = func() {
		group = group / (n - len(result)) //重要的分组条件 例如  n==4 group==24  但是result=0 所以就是6来进行递归找
		for i := 1; i <= n; i++ {
			if visited[i] {
				continue
			}
			if k > group {
				k = k - group //这里边主要找到第几个数字是起始位置
				continue
			}
			visited[i] = true
			result = append(result, strconv.Itoa(i))
			dfs()
		}
	}
	dfs()

	res := ""
	for i := 0; i < len(result); i++ {
		res = res + result[i]
	}
	return res
}

// lc-78 nums数组不重复 要求返回全部子集 并且不包含重复元素
// complete sunsets
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	var sub func(index int)
	list, emptySet := make([]int, 0), make([]int, 0)
	result = append(result, emptySet)
	sub = func(index int) {
		//递归退出条件
		if index == len(nums) {
			return
		}
		for i := index; i < len(nums); i++ {
			list = append(list, nums[i])
			temp := make([]int, len(list))
			copy(temp, list)
			result = append(result, temp)
			sub(i + 1) //这里边是sub(i+1) 还是sub(index+1)呢 注意这里边措辞呗
			list = list[0 : len(list)-1]
		}
	}
	sub(0)
	return result
}
