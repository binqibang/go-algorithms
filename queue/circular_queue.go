package queue

// MyCircularQueue
// @author: binqibang
// @create: 2022-08-02 10:49
type MyCircularQueue struct {
	elems []int
	front int
	rear  int
}

func CircularQueue(k int) MyCircularQueue {
	elems := make([]int, k+1)
	return MyCircularQueue{elems, 0, 0}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.elems[q.front] = value
	q.rear = (q.rear + 1) % len(q.elems)
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % len(q.elems)
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elems[q.front]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	idx := (q.rear + len(q.elems) - 1) % len(q.elems)
	return q.elems[idx]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *MyCircularQueue) IsFull() bool {
	return (q.rear+1)%len(q.elems) == q.front
}
