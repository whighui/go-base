//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链
//表节点，返回 反转后的链表 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//
//
// 示例 2：
//
//
//输入：head = [5], left = 1, right = 1
//输出：[5]
//
//
//
//
// 提示：
//
//
// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
//
//
//
// 进阶： 你可以使用一趟扫描完成反转吗？
//
// Related Topics 链表 👍 1604 👎 0

package doc

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 反转链表 从中间部位呢呗
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left >= right {
		return head
	}
	count := right - left + 1
	result := &ListNode{}
	result.Next = head
	pre := result
	for pre != nil && left > 1 {
		pre = pre.Next
		left--
	}

	//最终尾节点
	tail := pre.Next
	//找到前置节点
	head = pre.Next

	for head != nil && count > 0 {
		next := head.Next
		head.Next = pre.Next
		pre.Next = head
		head = next
		count--
	}
	tail.Next = head

	return result.Next
}

//leetcode submit region end(Prohibit modification and deletion)
