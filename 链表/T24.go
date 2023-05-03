package 链表

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next
		tmp1 := cur.Next.Next.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = tmp
		cur.Next.Next.Next = tmp1

		cur = cur.Next.Next
	}
	return dummyHead.Next
}
