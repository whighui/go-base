package lc

/**
主要实现大部分回文串先关实现呢
*/

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

// 判断回文数量
func judgePalindromeNums() {

}
