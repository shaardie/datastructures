package datastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleLinkedListFromListAndBack(t *testing.T) {
	tests := []struct {
		name string
		list []interface{}
	}{
		{"empty", []interface{}{}},
		{"single", []interface{}{0}},
		{"multiple", []interface{}{0, 1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running %v\n", tt.name)
			dll := NewDoubleLinkedListFromList(tt.list)
			t.Log("asdasd")
			assert.Equal(t, tt.list, dll.toList())
		})
	}
}

func TestDoubleLinkedListString(t *testing.T) {
	tests := []struct {
		name string
		list []interface{}
	}{
		{"empty", []interface{}{}},
		{"single", []interface{}{0}},
		{"multiple", []interface{}{0, 1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running %v\n", tt.name)
			assert.Equal(t,
				fmt.Sprintf("%v", tt.list),
				fmt.Sprintf("%v", NewDoubleLinkedListFromList(tt.list)))
		})
	}
}

func TestDoubleLinkedListInsert(t *testing.T) {
	tests := []struct {
		name    string
		index   int
		element interface{}
		list    []interface{}
		result  []interface{}
	}{
		{
			"empty",
			0,
			1,
			[]interface{}{},
			[]interface{}{1},
		},
		{
			"change first",
			0,
			1,
			[]interface{}{2, 3, 4, 5},
			[]interface{}{1, 2, 3, 4, 5},
		},
		{
			"change somewhere",
			2,
			3,
			[]interface{}{1, 2, 4, 5},
			[]interface{}{1, 2, 3, 4, 5},
		},
		{
			"change last",
			3,
			4,
			[]interface{}{1, 2, 3, 5},
			[]interface{}{1, 2, 3, 4, 5},
		},
		{
			"change after last",
			4,
			5,
			[]interface{}{1, 2, 3, 4},
			[]interface{}{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running %v\n", tt.name)
			sll := NewDoubleLinkedListFromList(tt.list)
			sll.Insert(tt.index, tt.element)
			t.Logf("New Double Linked List of size %v: %v", sll.Length(), sll)
			assert.Equal(t, tt.result, sll.toList())
		})
	}
}

func TestDoubleLinkedList_Pop(t *testing.T) {
	tests := []struct {
		name          string
		index         int
		result        interface{}
		list          []interface{}
		resultingList []interface{}
	}{
		{
			"pop first",
			0,
			1,
			[]interface{}{1, 2, 3, 4, 5},
			[]interface{}{2, 3, 4, 5},
		},
		{
			"pop somewhere",
			2,
			3,
			[]interface{}{1, 2, 3, 4, 5},
			[]interface{}{1, 2, 4, 5},
		},
		{
			"pop last",
			4,
			5,
			[]interface{}{1, 2, 3, 4, 5},
			[]interface{}{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running %v\n", tt.name)
			sll := NewDoubleLinkedListFromList(tt.list)
			r := sll.Pop(tt.index)
			t.Logf("New Single Linked List of size %v: %v", sll.Length(), sll)
			assert.Equal(t, tt.result, r)
			assert.Equal(t, tt.resultingList, sll.toList())
		})
	}
}
