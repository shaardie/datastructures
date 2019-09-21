package datastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleLinkedListFromListAndBack(t *testing.T) {
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
			assert.Equal(t, tt.list, NewSingleLinkedListFromList(tt.list).toList())
		})
	}
}

func TestSingleLinkedListString(t *testing.T) {
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
				fmt.Sprintf("%v", NewSingleLinkedListFromList(tt.list)))
		})
	}
}

func TestSingleLinkedListInsert(t *testing.T) {
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
			4,
			5,
			[]interface{}{1, 2, 3, 4},
			[]interface{}{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running %v\n", tt.name)
			sll := NewSingleLinkedListFromList(tt.list)
			sll.Insert(tt.index, tt.element)
			t.Logf("New Single Linked List of size %v: %v", sll.Length(), sll)
			assert.Equal(t, tt.result, sll.toList())
		})
	}
}

func TestSingleLinkedList_Pop(t *testing.T) {
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
			sll := NewSingleLinkedListFromList(tt.list)
			r := sll.Pop(tt.index)
			t.Logf("New Single Linked List of size %v: %v", sll.Length(), sll)
			assert.Equal(t, tt.result, r)
			assert.Equal(t, tt.resultingList, sll.toList())
		})
	}
}
