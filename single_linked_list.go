package datastructures

import "fmt"

// SingleLinkedList is a basic implementation of a single linked list.
//
// For more information about this data structure take a look at
// https://en.wikipedia.org/wiki/Linked_list.
//
// There are no safeguards in this implementation, if you try to read or write
// after the length is the behaviour undefined.
type SingleLinkedList struct {
	first  *singleNode
	length int
}

type singleNode struct {
	value interface{}
	next  *singleNode
}

// String is the string representation of the single linked list.
// It is analog to the representation of a regular array.
func (sll *SingleLinkedList) String() string {
	// Result string
	var s string

	// First node
	n := sll.first

	// Loop through list
	for {
		// Empty list
		if n == nil {
			break
		}

		if s == "" {
			// Prints first element
			s = fmt.Sprintf("%v", n.value)
		} else {
			// Add more elements
			s = fmt.Sprintf("%v %v", s, n.value)
		}

		// Next node
		n = n.next
	}

	// Add surrounding brackets
	return fmt.Sprintf("[%v]", s)
}

// NewSingleLinkedListFromList creates a new single linked list from an array.
func NewSingleLinkedListFromList(list []interface{}) *SingleLinkedList {
	// Get length of the list
	length := len(list)

	// List empty
	if length == 0 {
		return &SingleLinkedList{}
	}

	// Set first node
	n := &singleNode{value: list[0]}
	sll := &SingleLinkedList{
		first:  n,
		length: length,
	}

	// Set the next node and move forward
	for index := 1; index < length; index++ {
		n.next = &singleNode{
			value: list[index],
		}
		n = n.next
	}
	return sll
}

// Length gives the length of the single linked list
func (sll *SingleLinkedList) Length() int {
	return sll.length
}

// Insert a new element into the single linked list.
//
// The element is inserted on the index `index` and all elements including the
// element previously on the index `index` are moved further backwards in the
// single linked list.
func (sll *SingleLinkedList) Insert(index int, element interface{}) {
	// Panic if index is smaller 0
	if index < 0 {
		panic("index < 0")
	}

	// Insert as first element
	if index == 0 {
		sll.first = &singleNode{value: element, next: sll.first}
		sll.length++
		return
	}

	// Get node before the place where the new node is added
	n := sll.getNode(index - 1)
	// New node
	newNode := &singleNode{value: element}
	// If there is a node after the index where to insert, attach it to the
	// new node
	if n.next != nil {
		newNode.next = n.next
	}
	// Insert new node
	n.next = newNode
	sll.length++
}

// Append a new element to the single linked list.
func (sll *SingleLinkedList) Append(element interface{}) {
	sll.Insert(sll.length, element)
}

// Prepend a new element to the single linked list.
func (sll *SingleLinkedList) Prepend(element interface{}) {
	sll.Insert(0, element)
}

// Get the element on the index `index`.
func (sll *SingleLinkedList) Get(index int) interface{} {
	return sll.getNode(index).value
}

// Pop the element on the index `index`.
func (sll *SingleLinkedList) Pop(index int) interface{} {
	// Panic if index is smaller 0
	if index < 0 {
		panic("index < 0")
	}

	// Pop first element
	if index == 0 {
		// Result
		v := sll.first.value
		// Remove first element
		sll.first = sll.first.next
		// Decrease length
		sll.length--
		return v
	}

	// Get node before the one to pop
	n := sll.getNode(index - 1)
	// Result
	v := n.next.value
	// Remove reference to remove element
	n.next = n.next.next
	// Decrease length
	sll.length--
	return v
}

// Delete the element on the index `index`.
func (sll *SingleLinkedList) Delete(index int) {
	sll.Pop(index)
}

// getNode get the node on the index `index`.
func (sll *SingleLinkedList) getNode(index int) *singleNode {
	// Get first node
	n := sll.first
	// Loop through nodes to find the one with index `index`
	for index > 0 {
		n = n.next
		index--
	}
	return n
}

// toList transforms SingleLinkedList to a regular array.
func (sll *SingleLinkedList) toList() []interface{} {
	// Create the array.
	list := make([]interface{}, sll.length)
	// Get first node
	n := sll.first
	// Loop through the nodes and add them to the array
	for index := 0; index < sll.length; index++ {
		list[index] = n.value
		n = n.next
	}
	return list
}
