package linkedlist

import (
	"fmt"
	"go-algorithm/set"
	"math"
)

// 有头链表结构体
type LinkedList struct {
	Head *LNode
	Len  int
}

type LNode struct {
	Next *LNode
	Data int
}

// 将相邻的两个节点翻转（交换位置）-1：
// 互换值法：互相交换data的值
func (ll *LinkedList) NeighborReverse() *LinkedList {
	if ll.Len < 2 {
		return nil
	}
	cur := ll.Head.Next
	//Next := cur.Next
	var d int
	for cur.Next != nil {
		d = cur.Data
		cur.Data = cur.Next.Data
		cur.Next.Data = d
		cur = cur.Next.Next
	}
	return ll
}

// 将相邻的两个节点翻转（交换位置）-2：
// 调整指针法：改变指针的指向实现交换
func (ll *LinkedList) NeighborReverse2() *LinkedList {
	if ll.Len < 2 {
		return nil
	}

	pre := ll.Head
	cur := pre.Next
	var tmp *LNode
	for cur != nil && cur.Next != nil {
		// 交换指针
		tmp = cur.Next.Next
		pre.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = tmp
		// 移位
		pre = cur
		cur = tmp
	}
	return ll
}



// 如果链表存在环，找出环的入口点，即找出末端节点指向了哪一个节点
func (ll *LinkedList) FindEntrance() *LNode {
	isLoop := ll.IsLoop()
	if isLoop == nil {
		return nil
	}

	// 推到过程：
	// 从第一个节点出发，当快、慢两个指针相遇时，假设慢指针走了s步，则快指针走了2s步，此时快指针领先n圈
	// 假设环长为r，则：s + nr = 2s, n >= 1
	// 所以有 s = nr（这里可以得出第一个结论：第一次相遇时，慢指针最多刚好走完一整个链表，因为此时n=2，当环入口为第一个节点时，r最大）
	// 假设 链表长为L，第一个节点到环入口节点的距离为 a，环入口节点到相遇节点为x，则有如下关系
	// a + x = (n - 1) * r + r
	// a + x = (n - 1) * r + (L - a)
	// a = (n - 1) * r + (L - a - x)
	// 第一次相遇，n=0，所以有：a = L - a - x
	// 先看L - a - x，这个表达式的结果就是：从相遇点继续往前走到环入口点的距离
	// 而这个距离刚好等于：从第一个节点到环入口点的距离 a
	// 所以，如果一个指针从第一个节点出发，第二个指针从第一个相遇点出发，每次走一步，则必定在入口点相遇

	f := ll.Head.Next
	s := isLoop
	for {
		f = f.Next
		s = s.Next
		if f == s {
			return f
		}
	}
}


// 判断链表是否有环
// 快慢指针法：快慢不一的两个指针同时从第一个节点出发，如果有环，则一定会相遇
func (ll *LinkedList) IsLoop() *LNode {
	// 同时从第一个节点出发
	slow := ll.Head.Next
	fast := ll.Head.Next
	for slow != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

// 对链表重新排序：如1->2->3->4->5  ==> 1->5->2->4->3
func (ll *LinkedList) ReOrder() *LinkedList {
	if ll.Len <= 2 {
		return ll
	}

	// 特征-1：链表前半段相对位置不变—— 1   2   3
	// 所以找到后半段的第一个节点
	mid := math.Ceil(float64(5) / 2)
	var i float64
	i = 1
	curr2 := ll.Head.Next
	for i < mid {
		curr2 = curr2.Next
		i++
	}

	newHead := curr2.Next
	curr2.Next = nil

	// 节点数量为0或1时，无需处理直接返回原链表
	//if sec == nil {
	//	return ll
	//}

	// 将后半段逆序

	var pre *LNode
	c := newHead
	var n *LNode
	for c != nil {
		n = c.Next   // 先记录下一个节点的位置，不然后面就没了
		c.Next = pre // 逆序：当前节点指向前一个节点
		// 继续遍历下一个节点，当前节点变为pre节点
		pre = c
		c = n
	}
	newHead = pre

	// 开始合并两条链表
	curr1 := ll.Head.Next
	var tmp *LNode
	for curr1 != nil {
		if newHead == nil {
			break
		}
		tmp = newHead.Next
		newHead.Next = curr1.Next
		curr1.Next = newHead

		curr1 = newHead.Next
		newHead = tmp
	}
	return ll
}

// 删除重复项-3：利用一个int32位图数据结构进行去重
// 时间复杂度：只需遍历一遍链表，为O(n)
// 空间复杂度：O(1)
func (ll *LinkedList) UniqueBitmap() *LinkedList {
	bitmap := set.NewBitMap(125)

	pre := ll.Head
	cur := ll.Head.Next
	for cur != nil {
		if bitmap.Has(cur.Data) {
			// 将节点踢出链表
			pre.Next = cur.Next
			cur = cur.Next
			continue
		}
		bitmap.Add(cur.Data)
		pre = cur
		cur = cur.Next
	}
	return ll
}

// 删除重复项-2：递归法
func (ll *LinkedList) UniqueRecursive() *LinkedList {
	// 时间复杂度：需要进行双重遍历，所以为O(n^2)，相对于顺序法，增加了许多额外的函数调用（即递归），效率更低
	// 空间复杂度：使用了三个变量，所以为O(1)
	ll.Head.Next, ll.Len = uniqueNode(ll.Head.Next)
	return ll
}

// 通过递归对值去重
func uniqueNode(node *LNode) (*LNode, int) {
	if node == nil {
		return node, 0
	}
	if node.Next == nil {
		return node, 1
	}
	var length int
	node.Next, length = uniqueNode(node.Next)
	length += 1 // 算上当前node节点

	pre := node      // 当前节点的前驱节点
	cur := node.Next // 当前节点
	for cur != nil {
		if cur.Data == node.Data {
			// 由于删除了cur节点，那原cur.next节点的前驱节点依然是当前的pre，不用改动
			length -= 1
			pre.Next = cur.Next
		} else {
			// 即将进入下一次循环，当前节点将变成前驱节点
			pre = cur
		}
		cur = cur.Next
	}
	return node, length
}

// 删除重复项-1：顺序删除
func (ll *LinkedList) Unique() *LinkedList {
	if ll.Len <= 1 {
		return ll
	}

	/**
	时间复杂度：使用了双重循环，所以为O(n^2)
	空间复杂度：整个过程只是额外使用了3个变量，无其他额外空间占用，所以为O(1)
	*/

	var outerCur *LNode // 外层循环的当前节点
	var innerCur *LNode // 内层循环的当前节点
	var innerPre *LNode // 内层循环的当前节点的上一个节点

	outerCur = ll.Head.Next

	for outerCur != nil {
		innerPre = outerCur
		innerCur = outerCur.Next
		for innerCur != nil {

			if innerCur.Data != outerCur.Data {
				innerPre = innerCur
				innerCur = innerCur.Next
				continue
			}
			// 相同则移除innerCurr
			// @todo 这里只是不再引用那个节点，但那个节点还是存在于内存中，是否需要手动释放节点的内存？
			innerPre.Next = innerCur.Next // 前一个节点直接跳过当前节点，而指向下一个节点，实现将当前节点踢出链表
			innerCur = innerCur.Next      // 继续比较下一个节点
		}
		outerCur = outerCur.Next
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
	firstNode := ll.Head.Next
	if firstNode == nil || firstNode.Next == nil {
		return ll
	}

	// 实例化两个变量，保存当前节点和下一个节点的地址引用
	var cur *LNode
	var next *LNode

	cur = firstNode.Next // 从第二个节点开始处理
	firstNode.Next = nil // 将原链表的第一个节点，设置为最后一个节点
	for cur != nil {
		next = cur.Next // 先将原链中的下一个节点保存起来，再不存就没有了

		// 将当前节点，插到头节点后面
		cur.Next = ll.Head.Next
		ll.Head.Next = cur
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
	firstNode := ll.Head.Next
	if firstNode == nil || firstNode.Next == nil {
		return ll
	}
	// 将原链表的第一个节点传递给函数处理，函数将返回逆序后的第一个节点
	// 头节点直接指向逆序后的第一个节点
	ll.Head.Next = reverseChild(ll.Head.Next)
	return ll
}

// 递归函数，完成链表的倒序
func reverseChild(node *LNode) *LNode {
	if nil == node || nil == node.Next {
		return node // 已经到底了，直接返回最后一个节点
	}

	// 这是已经完成逆序的子链表的第一个节点
	newHead := reverseChild(node.Next)

	// 经过上面处理后，当前节点的子节点，已经是子链表中的最后一个节点
	// 所以这里要将当前节点添加到子链的最后，完成当前链的逆序
	// e.g. 假设当前节点为2 ，当前链为2 -> 3 -> 4
	// 		经过5上面的递归调用后，变成了：2 -> 3 <- 4
	//		所以这里要将3指向2，2指向nil，即2成为了最后一个节点：2 <- 3 <- 4
	node.Next.Next = node
	node.Next = nil
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
	firstNode := ll.Head.Next
	if firstNode == nil || firstNode.Next == nil {
		return ll
	}
	// 声明两个变量，用来保存当前节点的前、后两个节点的指针
	var pre *LNode
	var next *LNode

	pre = firstNode       // 对于第二个节点来说，pre就是第一个节点
	cur := firstNode.Next // 从第二个节点开始遍历
	firstNode.Next = nil  // 第一个节点，逆序后将变成最后一个节点，所以next指针要设为nil
	for nil != cur {
		next = cur.Next // 先把下一个节点的指针保存好，不保存等下就没啦
		cur.Next = pre  // 开始逆序，将next指针，指向前一个节点
		pre = cur       // 当前节点，即为下一个节点的pre
		cur = next      // 继续处理下一个
	}
	ll.Head.Next = pre // 最后，将头节点指向最后一个节点，完成逆序

	return ll
}

// 初始化一个空的单向链表
func NewList() *LinkedList {
	n := &LinkedList{
		Head: &LNode{Next: nil, Data: 0},
		Len:  0,
	}
	return n
}

// 添加一个元素
func (ll *LinkedList) Add(a int) *LNode {
	/**
	 * 时间复杂度：O(n)
	 * 空间复杂度：O(1)
	 */
	newNode := &LNode{Next: nil, Data: a}
	cur := ll.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	ll.Len += 1
	return newNode
}

// 打印所有元素
func (ll *LinkedList) PrintAllNode() {
	/**
	 * 时间复杂度：O(n)
	 * 空间复杂度：O(1)
	 */
	fmt.Printf("Len %d(", ll.Len)
	cur := ll.Head.Next
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
	fmt.Print(")")
}

// 删除节点
// 时间复杂度：删除非末端节点为O(1)，删除末端节点为O(n)，平均复杂度为O(1)
// 空间复杂度：最多需要2个变量存放指针，为O(1)
func (ll *LinkedList) Del(node *LNode) *LinkedList {
	if node == nil || ll.Len == 0 {
		return ll
	}

	n1 := node.Next // node节点的下一个节点
	if n1 != nil {
		// 因为不知道node节点的前一个节点在哪里，所以无法直接删除node
		// 但是可以知道node节点的下一个节点，删除下一个节点是O(1)
		// 在删除下一个节点前，将其内容复制到node，最终看起来就像是删除了node，实际上删除的是node 的next结点
		node.Next = n1.Next
		node.Data = n1.Data
		ll.Len -= 1
		// 最后取消n1节点的引用，让GC回收内存
		n1 = nil
		return ll
	}

	// 如果是最后一个节点，则必须从头遍历链表，直到找到前一个节点
	cur := ll.Head
	for cur.Next != node {
		cur = cur.Next
	}
	cur.Next = nil
	ll.Len -= 1
	return ll
}
