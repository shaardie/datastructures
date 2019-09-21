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

func (sll *SingleLinkedList) Length() int {
	return sll.length
}

func (sll *SingleLinkedList) Insert(index int, element interface{}) {

	switch {
	case index < 0:
	case index == 0:
	}

	if index < 0 {
		panic("index < 0")
	}

	if index == 0 {
		sll.first = &node{value: element, next: sll.first.next}
		sll.length++
		return
	}

	n := sll.getNode(index - 1)
	sll.insertAfter(&node{value: element}, n)
	sll.length++
}

func (sll *SingleLinkedList) Append(element interface{}) {
	sll.Insert(sll.Length(), element)
}

func (sll *SingleLinkedList) Prepend(element interface{}) {
	sll.Insert(0, element)
}

func (sll *SingleLinkedList) Get(index int) interface{} {
	return sll.getNode(index).value
}

func (sll *SingleLinkedList) Pop(index int) interface{} {

	if index < 0 {
		panic("index < 0")
	} else if index == 0 {
		v := sll.first.value
		sll.first = sll.first.next
		sll.length--
		return v
	}

	n := sll.first
	index--
	for index != 0 {
		n = n.next
		index--
	}

	v := n.next.value
	n.next = n.next.next
	sll.length--

	return v
}

func (sll *SingleLinkedList) Delete(index int) {
	sll.Pop(index)
}

func (sll *SingleLinkedList) insertAfter(first *node, after *node) {
	if after != nil {
		after.next = first.next
	}
	first.next = after
}

func (sll *SingleLinkedList) getNode(index int) *node {
	n := sll.first
	for index > 0 {
		n = n.next
		index--
	}
	return n
}
