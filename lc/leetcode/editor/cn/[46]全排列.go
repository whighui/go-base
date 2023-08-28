//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
//
// Related Topics 数组 回溯 👍 2630 👎 0

package doc

// leetcode submit region begin(Prohibit modification and deletion)

// lc-46  全排列
// 时间复杂度O(n * n!)  O(N)进行数组拷贝 O(N!)需要遍历这些
func permute(nums []int) [][]int {
	if len(nums) == 0 || nums == nil {
		return nil
	}
	result, list, visited := make([][]int, 0), make([]int, 0), make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(list) == len(nums) {
			temp := make([]int, len(list)) //这里边就是Temp需要指定长度呢 如果不指定长度就是默认为0  就是不发生拷贝
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
			dfs()
			visited[i] = false
			list = list[:len(list)-1] //这里边就是去除上一个元素呗 因为这里边就是全排列 就是跟全组合并不是一样的呗
		}
	}
	dfs()
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
