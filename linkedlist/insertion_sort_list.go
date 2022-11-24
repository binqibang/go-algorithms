package linkedlist

// InsertionSortList
// @summary: LeetCode #147 (Medium)
// @author : binqibang
// @created: 2022/11/24 17:19
func InsertionSortList(head *ListNode) *ListNode {
	// 在头结点前添加辅助结点，便于在头节点前插入
	dummyHead := &ListNode{Val: 0}
	dummyHead.Next = head
	// 维护 lastSorted 为已排序最后结点
	lastSorted, curr := head, head.Next
	for curr != nil {
		if lastSorted.Val <= curr.Val {
			lastSorted = curr
		} else {
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			// lastSorted 指向下一个待排序结点
			lastSorted.Next = curr.Next
			// 插入 curr
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}
