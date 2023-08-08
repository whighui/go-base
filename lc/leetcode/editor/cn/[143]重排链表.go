//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//
//L0 → L1 → … → Ln - 1 → Ln
//
//
// 请将其重新排列后变为：
//
//
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [1,2,3,4]
//输出：[1,4,2,3]
//
// 示例 2：
//
//
//
//
//输入：head = [1,2,3,4,5]
//输出：[1,5,2,4,3]
//
//
//
// 提示：
//
//
// 链表的长度范围为 [1, 5 * 10⁴]
// 1 <= node.val <= 1000
//
//
// Related Topics 栈 递归 链表 双指针 👍 1337 👎 0

package doc
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//lc-143 重排链表
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//思路找到中间节点 反转后边 进行连接
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	//把mid之前的节点找到呗
	mid := findMiddleNode(head)
	head1 := reverseListNode(mid.Next)
	mid.Next = nil

	for head1 != nil {
		next, next1 := head.Next, head1.Next
		head.Next = head1
		head1.Next = next
		head, head1 = next, next1
	}
}

//寻找中间节点 偶数找前边 奇数找中间
func findMiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseListNode(head *ListNode) *ListNode {
	var cur *ListNode
	for head != nil {
		next := head.Next
		head.Next = cur
		cur = head
		head = next
	}
	return cur
}



//leetcode submit region end(Prohibit modification and deletion)
