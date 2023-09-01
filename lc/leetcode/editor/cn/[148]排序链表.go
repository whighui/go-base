//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//
//
// 示例 2：
//
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
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
// 链表中节点的数目在范围 [0, 5 * 10⁴] 内
// -10⁵ <= Node.val <= 10⁵
//
//
//
//
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
//
// Related Topics 链表 双指针 分治 排序 归并排序 👍 2086 👎 0

package doc

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//还没有提交 需要进行提交呢
// lc-148 排序链表  题目要求时间复杂度为O(nlogn) 那么就是只有堆排序、快速排序、归并排序
// 由于快速排序并不是稳定排序 归并排序属于稳定排序 所以在这里就是使用归并排序来做呢
//链表进行排序呗  时间复杂度要求O(NlogN) 在这里就是主要意思是啥呢 要求比较快的排序算法呗 归并排序啥的
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return mergeSort(head)
}

// 先进行分裂链表呗
func mergeSort(head *ListNode) *ListNode {
	if head.Next == nil { //递归只升到最后一个元素
		return head
	}
	//寻找到中间节点 奇数正常 偶数找到对称之后的那个节点
	var slowPre *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil { //寻找到偶数前边一个节点
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	slowPre.Next = nil

	//上述操作是为了把链表节点都进行断开
	//下边操作是为了把断开的链表进行合并
	l1 := mergeSort(head)
	l2 := mergeSort(slow)
	return merge(l1, l2)
}

// merge进行链表排序
func merge(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	res := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			res.Next = l1
			l1 = l1.Next
		} else {
			res.Next = l2
			l2 = l2.Next
		}
		res = res.Next
	}
	if l1 == nil {
		res.Next=l2
	}
	if l2 == nil {
		res.Next=l1
	}

	return result.Next
}

//leetcode submit region end(Prohibit modification and deletion)
