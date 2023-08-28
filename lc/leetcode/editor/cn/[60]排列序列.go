//给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
//
// 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
//
//
// "123"
// "132"
// "213"
// "231"
// "312"
// "321"
//
//
// 给定 n 和 k，返回第 k 个排列。
//
//
//
// 示例 1：
//
//
//输入：n = 3, k = 3
//输出："213"
//
//
// 示例 2：
//
//
//输入：n = 4, k = 9
//输出："2314"
//
//
// 示例 3：
//
//
//输入：n = 3, k = 1
//输出："123"
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
// 1 <= k <= n!
//
//
// Related Topics 递归 数学 👍 791 👎 0

package doc

import "strconv"

//返回第K个排列呗 在这里哈呢需要点什么手段呢

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
	result, visited := make([]string, 0), make([]bool, n+1)
	var dfs func()
	dfs = func() {
		//递归退出条件就是没有写呗 在这里哈呢 就是要指定递归退出条件呗
		if len(result) == n {
			return
		}
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
			//visited[i] = false   不要有visited[i]==false这个条件 正常我们只需要拜访一圈呗
		}
	}
	dfs()

	res := ""
	for i := 0; i < len(result); i++ {
		res = res + result[i]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
