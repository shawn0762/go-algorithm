package main

import "fmt"

func main() {
	l := newList()
	//l.add(2).add(5).add(3).printAllNode()
	//l.add(2).add(5).add(3).reverse_direct().printAllNode()
	//l.add(2).add(5).add(3).reverse_recursive().printAllNode()
	l.add(2).add(5).add(3).reverse_insert().printAllNode()
}

type LinkedList struct {
	head *LNode
	len  int
}

type LNode struct {
	next *LNode
	data int
}

/**
	链表倒序-3：插入法
	遍历链表，从第二个节点开始，逐个将节点插入到头节点后面
	时间复杂度：整个过程只需要遍历一次链表，所以为O(n)，n为链表节点数量
	空间复杂度：相比递归，节省了栈操作的消耗；相比就地逆序，节省了一个pre变量
 */
func (ll *LinkedList) reverse_insert() *LinkedList {
	// 空链表，或链表只有一个结点，则直接返回，无需处理
	firstNode := ll.head.next
	if firstNode == nil || firstNode.next == nil {
		return ll
	}

	// 实例化两个变量，保存当前节点和下一个节点的地址引用
	var cur *LNode
	var next *LNode

	cur = firstNode.next // 从第二个节点开始处理
	firstNode.next = nil // 将原链表的第一个节点，设置为最后一个节点
	for cur != nil {
		next = cur.next // 先将原链中的下一个节点保存起来，再不存就没有了

		// 将当前节点，插到头节点后面
		cur.next = ll.head.next
		ll.head.next = cur
		// 继续处理下一个节点
		cur = next
	}
	return ll
}


/**
 * 链表倒序-2：递归实现逆序
 * 利用递归先完成子链表的倒序，然后将当前节点追加到子链的末端，最后将头节点指向逆序后的第一个节点，完成倒序
 *
 * 时间复杂度：整个过程只需要遍历一次链表，所以为O(n)，n为链表节点数量
 * 空间复杂度：相比就地逆序，节省了两个变量，但递归调用需要额外的压栈和弹栈操作，相比之下性能会有所下降
 */
func (ll *LinkedList) reverse_recursive() *LinkedList {
	// 空链表，或链表只有一个结点，则直接返回，无需处理
	firstNode := ll.head.next
	if firstNode == nil || firstNode.next == nil {
		return ll
	}
	// 将原链表的第一个节点传递给函数处理，函数将返回逆序后的第一个节点
	// 头节点直接指向逆序后的第一个节点
	ll.head.next = reverseChild(ll.head.next)
	return ll
}

// 递归函数，完成链表的倒序
func reverseChild(node *LNode) *LNode {
	if nil == node || nil == node.next {
		return node // 已经到底了，直接返回最后一个节点
	}

	// 这是已经完成逆序的子链表的第一个节点
	newHead := reverseChild(node.next)

	// 经过上面处理后，当前节点的子节点，已经是子链表中的最后一个节点
	// 所以这里要将当前节点添加到子链的最后，完成当前链的逆序
	// e.g. 假设当前节点为2 ，当前链为2 -> 3 -> 4
	// 		经过5上面的递归调用后，变成了：2 -> 3 <- 4
	//		所以这里要将3指向2，2指向nil，即2成为了最后一个节点：2 <- 3 <- 4
	node.next.next = node
	node.next = nil
	return newHead
}

/**
 * 链表倒序-1：就地逆序
 * 从第二个节点开始遍历，将next指针修改为前一个节点，最后将头节点指向最后遍历的那个节点
 * 时间复杂度：整个过程只需要遍历一次链表，所以为O(n)，n为链表节点数量
 * 空间复杂度：无论链表节点数量是多少，整个过程都只需要两个变量：pre和next，用来保存前后指针，所以为O(1)
 */
func (ll *LinkedList) reverse_direct() *LinkedList {
	// 空链表，或链表只有一个结点，则直接返回，无需处理
	firstNode := ll.head.next
	if firstNode == nil || firstNode.next == nil {
		return ll
	}
	// 声明两个变量，用来保存当前节点的前、后两个节点的指针
	var pre *LNode
	var next *LNode

	pre = firstNode       // 对于第二个节点来说，pre就是第一个节点
	cur := firstNode.next // 从第二个节点开始遍历
	firstNode.next = nil  // 第一个节点，逆序后将变成最后一个节点，所以next指针要设为nil
	for nil != cur {
		next = cur.next // 先把下一个节点的指针保存好，不保存等下就没啦
		cur.next = pre  // 开始逆序，将next指针，指向前一个节点
		pre = cur       // 当前节点，即为下一个节点的pre
		cur = next      // 继续处理下一个
	}
	ll.head.next = pre // 最后，将头节点指向最后一个节点，完成逆序

	return ll
}

// 初始化一个空的单向链表
func newList() *LinkedList {
	n := &LinkedList{
		head: &LNode{next: nil, data: 0},
		len:  0,
	}
	return n
}

// 添加元素的方法
func (ll *LinkedList) add(a int) *LinkedList {
	ll.len += 1
	newNode := &LNode{next: nil, data: a}
	cur := ll.head
	next := cur.next
	for next != nil {
		cur = cur.next
		next = cur.next
	}
	cur.next = newNode
	return ll
}

// 打印所有元素
func (ll *LinkedList) printAllNode() {
	cur := ll.head.next
	fmt.Printf("Len %d(", ll.len)
	for cur != nil {
		fmt.Print(cur.data, " ")
		cur = cur.next
	}
	fmt.Print(")")
}
