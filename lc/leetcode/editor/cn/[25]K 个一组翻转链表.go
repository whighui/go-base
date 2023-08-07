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
//K个一组反转链表呢
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	tempHead := head
	reverseCnt, group := 0, k
	for tempHead != nil && group > 0 {
		tempHead = tempHead.Next
		group--
		if group == 0 {
			reverseCnt++
			group = k
		}
	}
	//开始进行反转 反转之前要记住头尾节点呢
	result, tail := &ListNode{}, &ListNode{}
	res := result
	for i := 0; i < reverseCnt; i++ {
		tail = head
		group = k
		var cur *ListNode
		for head != nil && group > 0 {
			next := head.Next
			head.Next = cur
			cur = head
			head = next
			group--
		}
		result.Next = cur
		//连接
		result = tail //重置result 找到反转之前头结点
	}
	tail.Next = head //最后链接没有进行反转的
	return res.Next
}

//leetcode submit region end(Prohibit modification and deletion)
