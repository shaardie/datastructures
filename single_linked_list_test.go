package datastructures

import (
	"testing"
)

func TestSingleLinkedList(t *testing.T) {

	list := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	sll := NewSingleLinkedList()

	for _, element := range list {
		sll.Append(element)
		t.Log(sll.Length())
	}

	t.Logf("List: %v", list)
	t.Logf("Single Linked List: %v", sll)
}
