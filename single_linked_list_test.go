package datastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleLinkedList(t *testing.T) {
	list := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	sll := NewSingleLinkedList()
	for index, element := range list {
		sll.Append(element)
		assert.Equal(t, index+1, sll.Length(), "Length wrong")
		t.Log(sll)
	}
	assert.Equal(t, fmt.Sprint(sll), fmt.Sprint(list))
	for index, element := range list {
		sll.Prepend(element)
		assert.Equal(t, index+len(list), sll.Length(), "Length wrong")
	}

	for i := 0; i < len(list); i++ {
		assert.Equal(t, list[i], sll.Pop(10))
	}

	for i := 0; i < len(list); i++ {
		assert.Equal(t, list[9-i], sll.Get(i))
	}

	for i := 0; i < len(list); i++ {
		sll.Delete(0)
		assert.Equal(t, 9-i, sll.Length())
	}

	assert.Panics(t, func() { sll.Pop(-1) })
}
