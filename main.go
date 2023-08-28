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

//链表进行排序呗  时间复杂度要求O(NlogN) 在这里就是主要意思是啥呢 要求比较快的排序算法呗 归并排序啥的
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return spiltList(head)
}

//先进行分裂链表呗
func spiltList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//寻找中间节点进行分裂呗
	slow, fast := head, head
	for fast != nil && fast.Next != nil { //寻找到偶数前边一个节点
		slow = slow.Next
		fast = fast.Next.Next
	}
	spiltList(slow)
	slow.Next = nil
	spiltList(head)
	return mergeList(head, slow)
}

//merge进行链表排序
func mergeList(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	res := result

	for l1 != nil || l2 != nil {
		val1, val2 := math.MinInt32, math.MinInt32
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		if val1 <= val2 {
			l1 = l1.Next
			res.Next = l1
		} else {
			l2 = l2.Next
			res.Next = l2
		}
		res = res.Next
	}
	return result.Next
}
