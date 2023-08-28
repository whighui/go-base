//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
//
// Related Topics 数组 回溯 👍 1421 👎 0

package doc

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)

// 全排列II 带有重复数字 所以这里边就是需要进行减枝被
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 || nums == nil {
		return nil
	}

	result, list, visited := make([][]int, 0), make([]int, 0), make([]bool, len(nums))
	var traceback func()
	sort.Ints(nums) //返回nums是以递增序列来返回呢
	traceback = func() {
		if len(list) == len(nums) { //因为本身不带索引，所以这里边就是直接以list数组长度来判断
			temp := make([]int, len(list))
			copy(temp, list)
			result = append(result, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			//解释这个 如果拜访过那就不在拜访 非常好理解
			if visited[i] {
				continue
			}

			//第一层 1（下标0），1（下标1），2    第一层for循环
			//第二层 1（下标1），1（下标0），2    第二层for循环 这一层理应该被剪支调
			//所以需要剪支前提条件 i>0 && nums[i]==nums[i-1]
			//并且nums[i]这个元素理应该遍历完了（第一层就遍历完了，在第二层遍历的时候应该去掉），
			//所以visited[i-1]===false  代表第一层遍历完事，如果第一层没有遍历完事 那么visited[i-1]=true
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}

			list = append(list, nums[i])
			visited[i] = true
			traceback() //这里边就是不传索引 就是永远不会发生索引越界呢
			visited[i] = false
			list = list[:len(list)-1]
		}
	}
	traceback()
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
