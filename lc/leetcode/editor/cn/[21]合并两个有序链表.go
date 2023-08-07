//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
//
// Related Topics 递归 链表 👍 3192 👎 0

package doc

// leetcode submit region begin(Prohibit modification and deletion)
type ListNode struct {
	Val  int
	Next *ListNode
}

//lc-21  合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	res := result
	for list1 != nil && list2 != nil {
		var val1, val2 int
		if list1 != nil {
			val1 = list1.Val
		}
		if list2 != nil {
			val2 = list2.Val
		}
		if val1 < val2 {
			result.Next = &ListNode{val1, nil}
			list1 = list1.Next
		} else {
			result.Next = &ListNode{val2, nil}
			list2 = list2.Next
		}
		result = result.Next
	}
	if list1 != nil {
		result.Next = list1
	}
	if list2 != nil {
		result.Next = list2
	}
	return res.Next
}

//leetcode submit region end(Prohibit modification and deletion)
