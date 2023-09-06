package lc

type ListNode struct {
	Val  int
	Next *ListNode
}

//-----------------------------------------------------------------------------------------------------------链表专题

// K个一组反转链表呢
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

// lc-21  合并两个有序链表
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
		if val1 > val2 {
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

// lc-143 重排链表
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// 思路找到中间节点 反转后边 进行连接
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	//把mid之前的节点找到呗
	mid := findMiddleNode(head)
	head1 := reverseListNode(mid.Next)
	mid.Next = nil

	for head != nil {
		next, next1 := head.Next, head1.Next
		head.Next = head1
		head1.Next = next
		head, head1 = next, next1
	}
}

// 寻找中间节点 偶数找前边 奇数找中间
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

// 找到链表中间节点 如果是偶数找到后边一个节点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
