package datastructures

import "fmt"

type SingleLinkedList struct {
	first  *node
	length int
}

type node struct {
	value interface{}
	next  *node
}

func (sll *SingleLinkedList) String() string {
	var s string
	n := sll.first
	for {
		if n == nil {
			break
		}
		if s == "" {
			s = fmt.Sprintf("%v", n.value)
		} else {
			s = fmt.Sprintf("%v %v", s, n.value)
		}
		n = n.next
	}
	return fmt.Sprintf("[%v]", s)
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		first:  nil,
		length: 0,
	}
}

func NewSingleLinkedListFromList(list []interface{}) *SingleLinkedList {

	// Get length of the list
	length := len(list)

	// List empty
	if length == 0 {
		return &SingleLinkedList{}
	}

	// Set first node
	n := &node{value: list[0]}
	sll := &SingleLinkedList{
		first:  n,
		length: length,
	}

	// Set the next node and move forward
	for index := 1; index < length; index++ {
		n.next = &node{
			value: list[index],
		}
		n = n.next
	}
	return sll
}

func (sll *SingleLinkedList) Length() int {
	return sll.length
}

func (sll *SingleLinkedList) Insert(index int, element interface{}) {

	if index < 0 {
		panic("index < 0")
	}

	if index == 0 {
		sll.first = &node{value: element, next: sll.first}
		sll.length++
		return
	}

	n := sll.getNode(index - 1)
	newNode := &node{value: element}
	if n.next != nil {
		newNode.next = n.next
	}
	n.next = newNode

	sll.length++
}

func (sll *SingleLinkedList) Append(element interface{}) {
	sll.Insert(sll.length, element)
}

func (sll *SingleLinkedList) Prepend(element interface{}) {
	sll.Insert(0, element)
}

func (sll *SingleLinkedList) Get(index int) interface{} {
	return sll.getNode(index).value
}

func (sll *SingleLinkedList) Pop(index int) interface{} {
	if index == 0 {
		v := sll.first.value
		sll.first = sll.first.next
		sll.length--
		return v
	}

	n := sll.getNode(index - 1)
	v := n.next.value
	n.next = n.next.next
	sll.length--
	return v
}

func (sll *SingleLinkedList) Delete(index int) {
	sll.Pop(index)
}

func (sll *SingleLinkedList) getNode(index int) *node {
	n := sll.first
	for index > 0 {
		n = n.next
		index--
	}
	return n
}

// Transform SingleLinkedList to a regular one
func (sll *SingleLinkedList) toList() []interface{} {
	list := make([]interface{}, sll.length)
	n := sll.first
	for index := 0; index < sll.length; index++ {
		list[index] = n.value
		n = n.next
	}
	return list
}
