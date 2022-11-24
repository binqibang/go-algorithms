package linkedlist

func addFirst(head *ListNode, val int) *ListNode {
	node := &ListNode{Val: val}
	node.Next = head
	return node
}

func CreateLinkedList(values []int) *ListNode {
	n := len(values)
	head := &ListNode{Val: values[n-1]}
	for i := n - 2; i >= 0; i-- {
		head = addFirst(head, values[i])
	}
	return head
}
