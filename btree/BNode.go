package btree

import (
	"fmt"
	"math"
)

// 树节点的结构
type BNode struct {
	Data       interface{}
	LeftChild  *BNode
	RightChild *BNode
}

// 判断输入的某个数组是否是某二元查找树的后序遍历的序列
// 假设二元查找树的顺序是：1234567，则其后序遍历的顺序是：1325764
func IsAfterOrder(arr []int, start int, end int) bool {
	// 特征-1：数组最后一个元素肯定是根节点
	rootData := arr[end]

	var i int

	// 特征-2：左子树所有节点的值都比根节点小，右子树所有节点的值都比根节点大
	for i = start; i < end; i++ {
		if arr[i] > rootData {
			break
		}
	}

	// 这样的话，i前面的值一定是小于rootData
	for j := start; j < i; j++ {
		if arr[j] >= rootData {
			return false
		}
	}

	rightMatch := true
	leftMatch := true

	if i <= end - 1 {
		rightMatch = IsAfterOrder(arr, i, end - 1)
	}
	if i - 1 >= start {
		leftMatch = IsAfterOrder(arr, start, i-1)
	}
	return leftMatch && rightMatch
}

// 将二元查找树转换成双向链表，要求不能创建任何新节点，只能调整节点的指向
var pHead *BNode
var pEnd *BNode

func Convert2linkedList(node *BNode) {
	if node == nil {
		return
	}

	// 先将左子树按二元查找树的顺序接入双向链表
	Convert2linkedList(node.LeftChild)
	node.LeftChild = pEnd
	if pEnd == nil {
		pHead = node
	} else {
		pEnd.RightChild = node
	}
	pEnd = node

	// 从这里开始，右子树会从最左下角按顺序接入双向链表的末端，最终形成完整的链表
	Convert2linkedList(node.RightChild)
}

// 判断两棵树是否相等
// 相等的条件：树结构一模一样，对应位置的值也相同
func IsEqual(root1 *BNode, root2 *BNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil {
		return false
	}
	if root1 != nil && root2 == nil {
		return false
	}
	if root1.Data == root2.Data {
		return IsEqual(root1.LeftChild, root2.LeftChild) && IsEqual(root1.RightChild, root2.RightChild)
	}
	return false
}

var maxSum int        // 最大和
var maxSumNode *BNode // 最大子树的根节点

// 找出和最大的子树
// 递归求出每一棵子树的和
func FindMaxSubTree(root *BNode) int {
	if root == nil {
		return 0
	}
	// 只支持数据类型为整形的树
	sum, ok := root.Data.(int)
	if !ok {
		panic("Type int allowed only")
	}

	l := FindMaxSubTree(root.LeftChild)
	r := FindMaxSubTree(root.RightChild)
	sum += l + r
	if sum > maxSum {
		maxSum = sum
		maxSumNode = root
	}

	return sum
}

func LevelPrint(root *BNode, level int) int {
	if root == nil || level < 0 {
		return 0
	} else if level == 0 {
		fmt.Print(root.Data)
		return 1
	} else {
		return LevelPrint(root.LeftChild, level-1) + LevelPrint(root.RightChild, level-1)
	}
}

// 将一个有序数组，转化为“二元查找树”
// s 起始下标
// @param e 结尾下标
func Arr2Tree(arr []int, s int, e int) *BNode {
	// 初始化一个空的根节点
	root := NewNode()

	if e < s {
		return nil
	}

	// 取中间节点
	mid := int(math.Ceil(float64(s+e) / 2))

	root.Data = arr[mid]
	root.LeftChild = Arr2Tree(arr, s, mid-1)
	root.RightChild = Arr2Tree(arr, mid+1, e)
	return root
}

func TreePrintln(bn *BNode) {
	// 先打印左边节点
	if bn.LeftChild != nil {
		TreePrintln(bn.LeftChild)
	}
	// 然后打印自身节点
	fmt.Println(bn.Data)
	// 最后打印右边节点
	if bn.RightChild != nil {
		TreePrintln(bn.RightChild)
	}
}

func NewNode() *BNode {
	return &BNode{}
}
