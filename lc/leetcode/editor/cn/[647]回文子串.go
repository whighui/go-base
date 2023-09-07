//给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
//
// 回文字符串 是正着读和倒过来读一样的字符串。
//
// 子字符串 是字符串中的由连续字符组成的一个序列。
//
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
//
//
// 示例 1：
//
//
//输入：s = "abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//
//
// 示例 2：
//
//
//输入：s = "aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 1230 👎 0

package doc

//跟最长回文子串很像

// leetcode submit region begin(Prohibit modification and deletion)
// lc-647 实现回文子串
// 输入：s = "aaa" 输出：6 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
func countSubstrings(s string) int {

	nums := 0
	var judgePalindromeNums func(i, j int)
	judgePalindromeNums = func(i, j int) {
		for i >= 0 && j < len(s) && s[i] == s[j] {
			i--
			j++
			nums++
		}
	}

	for index := range s {
		judgePalindromeNums(index, index)   // 判断奇数节点
		judgePalindromeNums(index, index+1) //判断偶数节点
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
