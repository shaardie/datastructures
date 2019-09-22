package datastructures

import "fmt"

type DoubleLinkedList struct {
	first  *doubleNode
	last   *doubleNode
	length int
}

type doubleNode struct {
	value    interface{}
	next     *doubleNode
	previous *doubleNode
}

// String is the string representation of the double linked list.
// It is analog to the representation of a regular array.
func (dll *DoubleLinkedList) String() string {
	// Result string
	var s string

	// First node
	n := dll.first

	// Loop through list
	for {
		// Empty list
		if n == nil {
			break
		}

		if n == dll.first {
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

// Length gives the length of the double linked list
func (dll *DoubleLinkedList) Length() int {
	return dll.length
}

func (dll *DoubleLinkedList) Insert(index int, element interface{}) {
	// Panic if index is smaller 0
	if index < 0 {
		panic("index < 0")
	}
	dll.length++
	if index == 0 {
		dll.first = &doubleNode{
			value: element,
			next:  dll.first,
		}
		return
	}
	if index == dll.length {
		dll.last = &doubleNode{
			value:    element,
			previous: dll.last,
		}
		return
	}
	n := dll.getNode(index)
	previous := n.previous
	newNode := &doubleNode{
		value:    element,
		previous: previous,
		next:     n,
	}
	previous.next = newNode
}

// NewSingleLinkedListFromList creates a new single linked list from an array.
func NewDoubleLinkedListFromList(list []interface{}) *DoubleLinkedList {
	// Get length of the list
	length := len(list)

	// List empty
	if length == 0 {
		return &DoubleLinkedList{}
	}

	// Set first node
	n := &doubleNode{value: list[0]}
	dll := &DoubleLinkedList{
		first:  n,
		length: length,
	}

	// Set the next node and the previous node and move forward
	for index := 1; index < length; index++ {
		n.next = &doubleNode{
			value:    list[index],
			previous: n,
		}
		n = n.next
	}

	// Set last node
	dll.last = n.next

	return dll
}

func (dll *DoubleLinkedList) toList() []interface{} {
	// Create the array.
	list := make([]interface{}, dll.length)
	// Get first node
	n := dll.first
	// Loop through the nodes and add them to the array
	for index := 0; index < dll.length; index++ {
		list[index] = n.value
		n = n.next
	}
	return list

}

func (dll *DoubleLinkedList) getNode(index int) *doubleNode {
	l := dll.length
	if index < l/2 {
		// Get first node
		n := dll.first
		// Loop through nodes to find the one with index `index`
		for index > 0 {
			n = n.next
			index--
		}
		return n
	}
	// Get last node
	n := dll.last
	// Loop through nodes to find the one with index `index`
	for index > 0 {
		n = n.previous
		index--
	}
	return n
}
