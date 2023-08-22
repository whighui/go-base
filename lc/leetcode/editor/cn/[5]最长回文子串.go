//给你一个字符串 s，找到 s 中最长的回文子串。
//
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 6698 👎 0
package doc

//leetcode submit region begin(Prohibit modification and deletion)

//最长回文子串呗 这种使用中心扩散法呗是最好的呗
func longestPalindrome(s string) string {

	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(i, i, s, res)
		res = maxPalindrome(i, i+1, s, res)
	}
	return res
}

//判断局部是否成对呗  在这里直接返回字符创就是最好的办法呗
func maxPalindrome(i, j int, s, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(sub) > len(res) {
		res = sub
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
