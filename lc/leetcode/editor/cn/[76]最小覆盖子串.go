//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//
//
// 注意：
//
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//
//
// 示例 1：
//
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
//
// 示例 2：
//
//
//输入：s = "a", t = "a"
//输出："a"
//解释：整个字符串 s 是最小覆盖子串。
//
//
// 示例 3:
//
//
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//
//
// 提示：
//
//
// m == s.length
// n == t.length
// 1 <= m, n <= 10⁵
// s 和 t 由英文字母组成
//
//
//
//进阶：你能设计一个在
//o(m+n) 时间内解决此问题的算法吗？
//
// Related Topics 哈希表 字符串 滑动窗口 👍 2595 👎 0

package doc

import "math"

// leetcode submit region begin(Prohibit modification and deletion)

// lc-76 最小覆盖子串
// s = "ADOBECODEBANC", t = "ABC"   输出："BANC"
func minWindow(s string, t string) string {
	window, need := make(map[byte]int), make(map[byte]int) //window是滑动窗口 need是给定字符串计算数量
	for i := 0; i < len(t); i++ {
		need[t[i]]++ //for range 是utf-8遍历方式 4个字节
		//for循环遍历 是unicode字符集遍历
	}
	left, right, validNum := 0, 0, 0
	leftBegin, rightEnd := 0, math.MaxInt //在遍历过程中找到最下长度呗
	for left <= right && right < len(s) {
		char := s[right]
		right++
		if _, ok := need[char]; ok {
			window[char]++
			if window[char] == need[char] {
				validNum++
			}
		}
		for validNum == len(need) { //缩短left指针
			if right-left < rightEnd-leftBegin { //这里边保存结果不对
				leftBegin = left
				rightEnd = right
			}
			char = s[left]
			left++
			//判断还是否符合条件呢
			if _, ok := need[char]; ok {
				window[char]--                 //need里边存在 window里边才会存在
				if need[char] > window[char] { //说明该字符数量少了
					validNum--
				}
			}
		}
	}
	if rightEnd == math.MaxInt {
		return ""
	}
	return s[leftBegin:rightEnd]
}

//leetcode submit region end(Prohibit modification and deletion)
