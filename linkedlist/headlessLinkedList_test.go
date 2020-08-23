package linkedlist

import "testing"

func TestHeadlessLinkedList_Add(t *testing.T) {
	ll := NewHeadlessLinkedList()
	//ll.add(4).add(9).add(1).printL()
	//ll.Add(4).Add(9).Add(1).Add(3).ReverseDirect().PrintL()
	ll.Add(4).Add(9).Add(1).Add(3).ReverseRecursive().PrintL()
}
