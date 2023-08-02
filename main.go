package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head5 := &ListNode{5, nil}
	head4 := &ListNode{4, head5}
	head3 := &ListNode{3, head4}
	head2 := &ListNode{2, head3}
	head := &ListNode{1, head2}
	reverseKGroup(head, 2)
}

// K个一组反转链表 末尾不足K个不进行反转
func reverseKGroup(head *ListNode, k int) *ListNode {
	result := &ListNode{}
	result.Next = head
	res := result
	reverseCnt, group := 0, k //表示反转频次
	for head != nil && group > 0 {
		group--
		head = head.Next
		if group == 0 {
			group = k
			reverseCnt++
		}
	}

	head = res.Next  //反转之前头结点
	tail := res.Next //反转之后尾节点
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
		res = tail
		tail = res.Next
	}
	res.Next = head
	//反转之前找到头尾节点
	return result.Next
}
