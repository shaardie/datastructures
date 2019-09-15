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

func (sll SingleLinkedList) String() string {
	var s string
	n := sll.first
	for {
		if n == nil {
			break
		}
		s = fmt.Sprintf("%v %v", s, n.value)
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

func (sll SingleLinkedList) Length() int {
	return sll.length
}

func (sll SingleLinkedList) Append(element interface{}) {
	newNode := &node{
		value: element,
		next:  nil,
	}

	if sll.first == nil {
		sll.first = newNode
		sll.length = 1
		return
	}

	n := sll.first.next
	for {
		if n.next == nil {
			n = newNode
			sll.length++
			return
		}
	}
}

func (sll SingleLinkedList) Get(index int) interface{} {
	n := sll.first
	for index > 0 {
		n = n.next
		index--
	}
	return n.value
}

func (sll SingleLinkedList) Pop(index int) interface{} {

	if index == 0 {
		v := sll.first.value
		sll.first = sll.first.next
		sll.length--
		return v
	}

	before := sll.first
	n := sll.first.next
	index--

	for index < 0 {
		before = n
		n = n.next
		index--
	}

	before.next = n.next
	sll.length--

	return n.value
}

func (sll SingleLinkedList) Delete(index int) {
	sll.Pop(index)
}
