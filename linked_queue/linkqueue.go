package linked_queue

import "go-algorithm/linkedlist"

// 使用链表实现一个队列
type Queue struct {
	List   *linkedlist.LinkedList
	Bottom *linkedlist.LNode
}

// 入队：将节点追加到队列底部
func (q *Queue) Push(node int) *Queue {
	q.Bottom = q.List.Add(node)
	return q
}

// 出队：弹出队列第一个节点，O(1)
func (q *Queue) Pop() *linkedlist.LNode {
	pop := q.List.Head.Next
	if pop == nil {
		panic("error, queue is empty")
	}
	q.List.Head.Next = pop.Next
	q.List.Len -= 1
	return pop
}

func (q *Queue) IsEmpty() bool {
	return q.List.Head.Next == nil
}

func (q *Queue) Size() int {
	return q.List.Len
}

func NewQueue() *Queue {
	list := linkedlist.NewList()
	return &Queue{
		List:   list,
		Bottom: list.Head,
	}
}
