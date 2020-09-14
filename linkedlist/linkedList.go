package linkedlist

import (
	"fmt"
	"go-algorithm/set"
	"math"
)

// 有头链表结构体
type LinkedList struct {
	head *LNode
	len  int
}

type LNode struct {
	next *LNode
	data int
}

// 对链表重新排序：如1->2->3->4->5  ==> 1->5->2->4->3
func (ll *LinkedList) ReOrder() *LinkedList {
	if ll.len <= 2 {
		return ll
	}

	// 特征-1：链表前半段相对位置不变—— 1   2   3
	// 所以找到后半段的第一个节点
	mid := math.Ceil(float64(5)/2)
	var i float64
	i = 1
	curr2 := ll.head.next
	for i < mid {
		curr2 = curr2.next
		i++
	}

	newHead := curr2.next
	curr2.next = nil


	// 节点数量为0或1时，无需处理直接返回原链表
	//if sec == nil {
	//	return ll
	//}

	// 将后半段逆序

	var pre *LNode
	c := newHead
	var n *LNode
	for c != nil {
		n = c.next // 先记录下一个节点的位置，不然后面就没了
		c.next = pre // 逆序：当前节点指向前一个节点
		// 继续遍历下一个节点，当前节点变为pre节点
		pre = c
		c = n
	}
	newHead = pre

	// 开始合并两条链表
	curr1 := ll.head.next
	var tmp *LNode
	for curr1 != nil {
		if newHead == nil {
			break
		}
		tmp = newHead.next
		newHead.next = curr1.next
		curr1.next = newHead

		curr1 = newHead.next
		newHead = tmp
	}
	return ll
}

// 删除重复项-3：利用一个int32位图数据结构进行去重
// 时间复杂度：只需遍历一遍链表，为O(n)
// 空间复杂度：O(1)
func (ll *LinkedList) UniqueBitmap() *LinkedList {
	bitmap := set.NewBitMap(125)

	pre := ll.head
	cur := ll.head.next
	for cur != nil {
		if bitmap.Has(cur.data) {
			// 将节点踢出链表
			pre.next = cur.next
			cur = cur.next
			continue
		}
		bitmap.Add(cur.data)
		pre = cur
		cur = cur.next
	}
	return ll
}

// 删除重复项-2：递归法
func (ll *LinkedList) UniqueRecursive() *LinkedList {
	// 时间复杂度：需要进行双重遍历，所以为O(n^2)，相对于顺序法，增加了许多额外的函数调用（即递归），效率更低
	// 空间复杂度：使用了三个变量，所以为O(1)
	ll.head.next, ll.len = uniqueNode(ll.head.next)
	return ll
}

// 通过递归对值去重
func uniqueNode(node *LNode) (*LNode, int) {
	if node == nil {
		return node, 0
	}
	if node.next == nil {
		return node, 1
	}
	var length int
	node.next, length = uniqueNode(node.next)
	length += 1 // 算上当前node节点

	pre := node      // 当前节点的前驱节点
	cur := node.next // 当前节点
	for cur != nil {
		if cur.data == node.data {
			// 由于删除了cur节点，那原cur.next节点的前驱节点依然是当前的pre，不用改动
			length -= 1
			pre.next = cur.next
		} else {
			// 即将进入下一次循环，当前节点将变成前驱节点
			pre = cur
		}
		cur = cur.next
	}
	return node, length
}

// 删除重复项-1：顺序删除
func (ll *LinkedList) Unique() *LinkedList {
	if ll.len <= 1 {
		return ll
	}

	/**
	时间复杂度：使用了双重循环，所以为O(n^2)
	空间复杂度：整个过程只是额外使用了3个变量，无其他额外空间占用，所以为O(1)
	*/

	var outerCur *LNode // 外层循环的当前节点
	var innerCur *LNode // 内层循环的当前节点
	var innerPre *LNode // 内层循环的当前节点的上一个节点

	outerCur = ll.head.next

	for outerCur != nil {
		innerPre = outerCur
		innerCur = outerCur.next
		for innerCur != nil {

			if innerCur.data != outerCur.data {
				innerPre = innerCur
				innerCur = innerCur.next
				continue
			}
			// 相同则移除innerCurr
			// @todo 这里只是不再引用那个节点，但那个节点还是存在于内存中，是否需要手动释放节点的内存？
			innerPre.next = innerCur.next // 前一个节点直接跳过当前节点，而指向下一个节点，实现将当前节点踢出链表
			innerCur = innerCur.next      // 继续比较下一个节点
		}
		outerCur = outerCur.next
	}
	return ll
}

/**
链表倒序-3：插入法
遍历链表，从第二个节点开始，逐个将节点插入到头节点后面
时间复杂度：整个过程只需要遍历一次链表，所以为O(n)，n为链表节点数量
空间复杂度：相比递归，节省了栈操作的消耗；相比就地逆序，节省了一个pre变量
*/
func (ll *LinkedList) ReverseInsert() *LinkedList {
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
func (ll *LinkedList) ReverseRecursive() *LinkedList {
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
func (ll *LinkedList) ReverseDirect() *LinkedList {
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
func NewList() *LinkedList {
	n := &LinkedList{
		head: &LNode{next: nil, data: 0},
		len:  0,
	}
	return n
}

// 添加一个元素
func (ll *LinkedList) Add(a int) *LNode {
	/**
	 * 时间复杂度：O(n)
	 * 空间复杂度：O(1)
	 */
	newNode := &LNode{next: nil, data: a}
	cur := ll.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode
	ll.len += 1
	return newNode
}

// 打印所有元素
func (ll *LinkedList) PrintAllNode() {
	/**
	 * 时间复杂度：O(n)
	 * 空间复杂度：O(1)
	 */
	fmt.Printf("Len %d(", ll.len)
	cur := ll.head.next
	for cur != nil {
		fmt.Print(cur.data, " ")
		cur = cur.next
	}
	fmt.Print(")")
}

// 删除节点
// 时间复杂度：删除非末端节点为O(1)，删除末端节点为O(n)，平均复杂度为O(1)
// 空间复杂度：最多需要2个变量存放指针，为O(1)
func (ll *LinkedList) Del(node *LNode) *LinkedList {
	if node == nil || ll.len == 0 {
		return ll
	}

	n1 := node.next // node节点的下一个节点
	if n1 != nil {
		// 因为不知道node节点的前一个节点在哪里，所以无法直接删除node
		// 但是可以知道node节点的下一个节点，删除下一个节点是O(1)
		// 在删除下一个节点前，将其内容复制到node，最终看起来就像是删除了node，实际上删除的是node 的next结点
		node.next = n1.next
		node.data = n1.data
		ll.len -= 1
		// 最后取消n1节点的引用，让GC回收内存
		n1 = nil
		return ll
	}

	// 如果是最后一个节点，则必须从头遍历链表，直到找到前一个节点
	cur := ll.head
	for cur.next != node {
		cur = cur.next
	}
	cur.next = nil
	ll.len -= 1
	return ll
}
