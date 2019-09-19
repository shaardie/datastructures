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
	l := 1
	for _, element := range list {
		sll.Append(element)
		assert.Equal(t, l, sll.Length(), "Length wrong")
		l++
	}
	assert.Equal(t, fmt.Sprint(sll), fmt.Sprint(list))
	for _, element := range list {
		sll.Prepend(element)
		assert.Equal(t, l, sll.Length(), "Length wrong")
		l++
	}

	for i := 0; i < len(list); i++ {
		assert.Equal(t, list[i], sll.Pop(10))
		t.Logf("%v", sll)
	}
}
