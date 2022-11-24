package linkedlist

import "testing"

func TestInsertionSortList(t *testing.T) {
	values := []int{-1, 5, 3, 4, 0}
	head := CreateLinkedList(values)
	sorted := InsertionSortList(head)
	t.Log(sorted)
}
