package lc

import "math"

/**
主要是实现双指针部分 双指针思想被
     * 1、我们在字符串S中使用双指针中的左右指针技巧，初始化left = right = 0，把索引左闭右开区间[left, right)称为一个「窗
     * 2、我们先不断地增加right指针扩大窗口[left, right)，直到窗口中的字符串符合要求（包含了T中的所有字符）。
     * 3、此时，我们停止增加right，转而不断增加left指针缩小窗口[left, right)，直到窗口中的字符串不再符合要求就不在进行移动             呗（不包含T中的所有字符了）。同时，每次增加left，我们都要更新一轮结果。
     * 4、重复第 2 和第 3 步，直到right到达字符串S的尽头
*/

// lc-76 最小覆盖子串
// s = "ADOBECODEBANC", t = "ABC"   输出："BANC"
func minWindow(s string, t string) string {
	window, need := make(map[byte]int), make(map[byte]int) //window是滑动窗口 need是给定字符串计算数量
	for i := 0; i < len(t); i++ {
		need[t[i]]++ //for range 是utf-8遍历方式 4个字节
		//for循环遍历 是unicode字符集遍历
	}
	left, right, validNum := 0, 0, 0
	leftBegin, rightEnd := 0, 0 //在遍历过程中找到最下长度呗

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
			if right-left < rightEnd-leftBegin {
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
	return s[leftBegin:rightEnd]
}

//滴滴面试题  类似于最小覆盖子串  只不过对其省略了一些条件
// 有一个整形数组大小为n
// 所有的元素都是12345 这五个元素在数组里是随机排列的
// 求最短子序列同时包含12345而且最短
// [1 3 3 2 4 2 4 3 5 3 5 5 3 1 4 4 2]

func minLength(nums []int) int {
	if len(nums) < 5 {
		return 0
	}
	window := make(map[int]int)
	left, right := 0, 0
	minLen := math.MaxInt
	for left <= right && right < len(nums) {
		window[nums[right]]++
		right++
		//这里边window可以不需要 排除不需要的字符呢
		for len(window) == 5 {
			if right-left < minLen {
				minLen = right - left
			}
			//记录最小长度
			window[nums[left]]--
			if window[nums[left]] == 0 {
				delete(window, nums[left])
			}
			left++
		}
	}
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}
