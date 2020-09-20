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

func TestLinkedList_ReOrder(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.ReOrder().PrintAllNode()
}

func TestLinkedList_IsLoop(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add(2)
	entry := l.Add(3) // 环形入口点
	l.Add(4)
	last := l.Add(5)

	//entry = nil
	last.Next = entry
	isLoop := l.IsLoop()
	// 第一次相遇的点肯定是4
	if isLoop == nil || isLoop.Data != 4 {
		println("is not loop")
		t.Error("unexpected result")
	}
}



func TestLinkedList_FindEntrance(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add(2)
	entry := l.Add(3) // 环形入口点
	l.Add(4)
	last := l.Add(5)
	last.Next = entry
	e := l.FindEntrance()
	if e == nil || e.Data != 3 {
		println("is not loop")
		t.Error("unexpected result")
	} else {
		println("Find circle entrace success: ", e.Data)
	}
}

func TestLinkedList_NeighborReverse(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.NeighborReverse().PrintAllNode()
}

func TestLinkedList_NeighborReverse2(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.NeighborReverse2().PrintAllNode()
}