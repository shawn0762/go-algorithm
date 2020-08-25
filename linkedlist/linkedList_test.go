package linkedlist

import "testing"

func TestLinkedList_Add(t *testing.T) {
		l := NewList()
		l.Add(2)
		l.Add(5)
		l.Add(3)
		l.PrintAllNode()
}

func TestLinkedList_Unique(t *testing.T) {
	l := NewList()
	l.Add(2)
	l.Add(5)
	l.Add(3)
	l.Add(5)
	l.Add(1)
	l.Add(2)
	l.Unique().PrintAllNode()
}

func TestLinkedList_Del(t *testing.T) {
	l := NewList()
	//l.Add(2)
	//l.Add(5)
	n := l.Add(3)
	l.PrintAllNode()

	l.Del(n)
	l.PrintAllNode()

}

func TestLinkedList_UniqueRecursive(t *testing.T) {
	l := NewList()
	l.Add(3)
	l.Add(2)
	l.Add(3)
	l.Add(5)
	l.UniqueRecursive().PrintAllNode()
}

func TestLinkedList_UniqueBitmap(t *testing.T) {
	l := NewList()
	l.Add(3)
	l.Add(2)
	l.Add(3)
	l.Add(5)
	l.Add(5)
	l.UniqueBitmap().PrintAllNode()
}