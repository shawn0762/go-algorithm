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

func convert2linkedList(node *BNode) *BNode {

	//newLeftRoot := convert2linkedList(node.LeftChild)
	//newRightRoot := convert2linkedList(node.LeftChild)


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
