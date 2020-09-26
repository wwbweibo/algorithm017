func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	node := &ListNode{}
	head = node
	for ; l1 != nil && l2 != nil; {
		if l1.Val >= l2.Val {
			node.Next = l2
			l2 = l2.Next
		} else {
			node.Next = l1
			l1 = l1.Next
		}
		node = node.Next
	}
	if l1 != nil{
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return head.Next
}