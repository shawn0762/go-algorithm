package linkedlist

import "fmt"

// 无头节点链表结构
type HeadlessLinkedList struct {
	first *LNode
	len   int
}

// 链表逆序-3：插入法逆序
// 无头链表，无法实现插入法逆序

// 链表逆序-2：递归逆序
// 先将当前节点的所有子节点逆序，然后将当前节点追加到末端
func (ll *HeadlessLinkedList) ReverseRecursive() *HeadlessLinkedList {
	if ll.len <= 1 {
		return ll
	}

	ll.first = reverseChildB(ll.first)
	return ll
}

// 通过递归调用，完成对子链的逆序
func reverseChildB(node *LNode) *LNode {
	if node == nil || node.next == nil {
		return node
	}

	newHead := reverseChildB(node.next)
	node.next.next = node
	node.next = nil
	return newHead
}

// 链表逆序-1：就地逆序
// 遍历全部节点，从第二个节点开始，将其指针指向前一个节点，最后将链表first属性指向最后一个节点，完成逆序
// 时间复杂度：O(n)，空间复杂度：O(1)
func (ll *HeadlessLinkedList) ReverseDirect() *HeadlessLinkedList {
	if ll.len <= 1 {
		return ll
	}

	// 需要三个变量，分别保存当前节点、前一个节点、后一个节点
	var cur *LNode
	var pre *LNode
	var next *LNode

	pre = ll.first
	cur = ll.first.next

	ll.first.next = nil // 将原链表中第一个节点，设为末端节点

	for cur != nil {
		next = cur.next
		cur.next = pre // 指向前一个节点
		// 继续处理下一个节点
		// 当前节点就是下一个节点的pre
		pre = cur
		cur = next
	}
	ll.first = pre
	return ll
}

// 初始化一个空链表
func NewHeadlessLinkedList() *HeadlessLinkedList {
	ll := &HeadlessLinkedList{first: nil, len: 0}
	return ll
}

// 添加一个新节点
func (ll *HeadlessLinkedList) Add(i int) *HeadlessLinkedList {
	// 新节点
	n := &LNode{next: nil, data: i}
	// 如果是空链表，则新节点为第一个节点
	if ll.len == 0 {
		ll.first = n
		ll.len = 1
		return ll
	}

	cur := ll.first
	for cur.next != nil {
		cur = cur.next
	}
	// 将新节点追加到链表末端
	cur.next = n
	ll.len += 1
	return ll
}

// 打印全部节点的值，
// 格式：Len 4(4 9 1 3)
func (ll *HeadlessLinkedList) PrintL() {
	cur := ll.first
	fmt.Printf("Len %d(", ll.len)
	for cur != nil {
		fmt.Print(cur.data, " ")
		cur = cur.next
	}
	fmt.Print(")")
}
