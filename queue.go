package datastructures

// Queue is a basic implementation of a queue.
//
// For more information the data structure take a look at
// https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
type Queue struct {
	dll *DoubleLinkedList
}

// Enqueue add and element to the queue.
func (q Queue) Enqueue(element interface{}) {
	q.dll.Append(element)
}

// Dequeue get the last element from the queue and removes it.
func (q Queue) Dequeue() interface{} {
	return q.dll.Pop(0)
}
