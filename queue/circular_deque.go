package queue

// MyCircularDeque
// @summary: LeetCode #641
// @author : binqibang
// @created: 2022/8/15 11:31
type MyCircularDeque struct {
	elems []int
	front int
	rear  int
	size  int
}

func CircularDeque(k int) MyCircularDeque {
	return MyCircularDeque{
		elems: make([]int, k+1),
		front: 0,
		rear:  0,
		size:  k + 1,
	}
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}
	q.front = (q.front + q.size - 1) % q.size
	q.elems[q.front] = value
	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}
	q.elems[q.rear] = value
	q.rear = (q.rear + 1) % q.size
	return true
}

func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % q.size
	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}
	q.rear = (q.rear + q.size - 1) % q.size
	return true
}

func (q *MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elems[q.front]
}

func (q *MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}
	idx := (q.rear + q.size - 1) % q.size
	return q.elems[idx]
}

func (q *MyCircularDeque) IsEmpty() bool {
	return q.front == q.rear
}

func (q *MyCircularDeque) IsFull() bool {
	return (q.rear+1)%q.size == q.front
}
