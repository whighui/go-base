package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(reverse(-123))
}
func reverse(x int) int {
	var res int64
	for x != 0 {
		tail := x % 10
		res = res*10 + int64(tail)
		x = x / 10
	}
	if res > math.MaxInt32 {
		return 0
	}
	if res < math.MinInt32 {
		return 0
	}
	return int(res)
}

// 最长回文子串呗 这种使用中心扩散法呗是最好的呗
func longestPalindrome(s string) string {

	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(i, i, s, res)
		res = maxPalindrome(i, i+1, s, res)
	}
	return res
}

// 判断局部是否成对呗  在这里直接返回字符创就是最好的办法呗
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
