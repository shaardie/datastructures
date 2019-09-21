package datastructures

// Stack is a basic implementation of a stack.
//
// For more information about the data structure take a look at
// https://en.wikipedia.org/wiki/Stack_(abstract_data_type).Stack
type Stack struct {
	sll *SingleLinkedList
}

// Push a new element onto the Stack.
func (s Stack) Push(element interface{}) {
	s.sll.Prepend(element)
}

// Pop the top element from the Stack.
func (s Stack) Pop() interface{} {
	return s.sll.Pop(0)
}
