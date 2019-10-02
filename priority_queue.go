package datastructures

import (
	"container/heap"
	"fmt"
)

// PriorityQueueItem is a single Element in the Priority Queue.
type PriorityQueueItem struct {
	// Value is the generic value of the element
	Value interface{}
	// Priority defines when this element is popped
	Priority int
	index    int
}

// String representation of the PriorityQueueItem
func (pqi PriorityQueueItem) String() string {
	return fmt.Sprintf("{%v %v}", pqi.Value, pqi.Value)
}

type priorityQueue []PriorityQueueItem

// PriorityQueue is a basic implementation of a Priority Queue.
//
// For more information about the data structure take a look at
// https://en.wikipedia.org/wiki/Priority_queue
type PriorityQueue struct {
	pq *priorityQueue
}

// NewPriorityQueue with items as optional initial PriorityQueueItems
func NewPriorityQueue(items ...PriorityQueueItem) *PriorityQueue {
	pq := make(priorityQueue, len(items))
	for i, item := range items {
		item.index = i
		pq[i] = item
	}
	heap.Init(&pq)
	return &PriorityQueue{&pq}
}

// String representation of the PriorityQueue
func (pq PriorityQueue) String() string {
	return fmt.Sprintf("%v", *pq.pq)
}

// Length returns the length of the PriorityQueue
func (pq PriorityQueue) Length() int {
	return pq.pq.Len()
}

// Pop the PriorityQueueItem with the highest priority
func (pq PriorityQueue) Pop() PriorityQueueItem {
	return heap.Pop(pq.pq).(PriorityQueueItem)
}

// Push a new PriorityQueueItem to the PrioristyQueue
func (pq PriorityQueue) Push(pqi PriorityQueueItem) {
	heap.Push(pq.pq, pqi)
}

// Len is the length
func (pq priorityQueue) Len() int { return len(pq) }

// Less compares two elements from index `i` and `j`
func (pq priorityQueue) Less(i, j int) bool {
	// Actually invert the comparator, since we want to have the one with
	// the highest priority.
	return pq[i].Priority > pq[j].Priority
}

// Swap the element from index `i` and `j`
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := pq.Len()
	item := x.(PriorityQueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
