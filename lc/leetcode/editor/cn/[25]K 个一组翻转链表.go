//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//
//
//提示：
//
//
// 链表中的节点数目为 n
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
//
//
//
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
//
//
//
//
// Related Topics 递归 链表 👍 2096 👎 0

package doc

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//K个一组反转链表 末尾不足K个不进行反转
func reverseKGroup(head *ListNode, k int) *ListNode {
	result := &ListNode{}
	res := result
	reverseCnt, group := 0, k //表示反转频次
	tempHead := head
	for head != nil && group > 0 {
		group--
		head = head.Next
		if group == 0 {
			group = k
			reverseCnt++
		}
	}

	head = tempHead  //反转之前头结点
	tail := tempHead //反转之后尾节点
	for i := 0; i < reverseCnt; i++ {
		group = k
		var cur *ListNode
		for head != nil && group > 0 {
			next := head.Next
			head.Next = cur
			cur = head
			head = next
			group--
		}
		res.Next = cur
		tail.Next = head //连接中间尾节点
		res = tail       //连接中间上一节点
		tail = head
	}
	res.Next = head
	//反转之前找到头尾节点
	return result.Next
}

//leetcode submit region end(Prohibit modification and deletion)
