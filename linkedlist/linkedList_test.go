package linkedlist

import "testing"

func TestLinkedList_Add(t *testing.T) {
		l := NewList()
		l.Add(2).Add(5).Add(3).PrintAllNode()
}