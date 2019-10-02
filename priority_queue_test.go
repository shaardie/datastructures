package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	items := []PriorityQueueItem{
		PriorityQueueItem{
			Value:    1,
			Priority: 1,
		},
		PriorityQueueItem{
			Value:    3,
			Priority: 3,
		},
	}

	pq := NewPriorityQueue(items...)
	pq.Push(PriorityQueueItem{Value: 4, Priority: 4})
	pq.Push(PriorityQueueItem{Value: 2, Priority: 2})
	pq.Push(PriorityQueueItem{Value: 0, Priority: 0})
	t.Logf("Current Priority Queue '%v'\n", pq)
	for i := pq.Length() - 1; i > 0; i-- {
		item := pq.Pop()
		t.Logf("Pop item '%v'\n", item)
		assert.Equal(t, i, item.Value)
		assert.Equal(t, i, item.Priority)
		t.Logf("Priority Queue after Pop '%v'\n", pq)
	}
}
