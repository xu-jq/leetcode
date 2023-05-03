package 链表

func reverseList(head *ListNode) *ListNode {
	temp := &ListNode{}
	slow := &ListNode{}
	fast := &ListNode{}
	slow = nil
	fast = head
	for fast != nil {
		temp = fast.Next
		fast.Next = slow
		slow = fast
		fast = temp
	}
	return slow
}
