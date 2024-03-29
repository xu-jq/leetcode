package 链表

func removeElements(head *ListNode, val int) *ListNode {
	node := &ListNode{}
	node.Next = head
	cur := node
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return node.Next
}
