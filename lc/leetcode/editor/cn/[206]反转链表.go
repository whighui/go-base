//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
// Related Topics 递归 链表 👍 3251 👎 0

package doc

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//还是for循环比较简单
func reverseList(head *ListNode) *ListNode {
	//这种递归的思想还是挺折磨人的呢
	if head == nil || head.Next == nil {
		return head
	}
	result := reverseList(head.Next) //result就是找到最末尾一个节点
	//此时程序往下走 result==5 head==4
	head.Next.Next = head //4->5
	head.Next = nil       //5->nil
	return result
}

//这里边分两种形式呗 递归反转和for循环反转
//leetcode submit region end(Prohibit modification and deletion)
