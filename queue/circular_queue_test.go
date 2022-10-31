package queue

import "testing"

func TestMyCircularQueue(t *testing.T) {
	myCircularQueue := CircularQueue(3)
	myCircularQueue.EnQueue(1) // return True
	myCircularQueue.EnQueue(2) // return True
	myCircularQueue.EnQueue(3) // return True
	myCircularQueue.EnQueue(4) // return False
	myCircularQueue.Rear()     // return 3
	myCircularQueue.IsFull()   // return True
	myCircularQueue.DeQueue()  // return True
	myCircularQueue.EnQueue(4) // return True
	myCircularQueue.Rear()     // return 4
}
