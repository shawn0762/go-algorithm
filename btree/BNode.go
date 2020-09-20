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

// 将一个有序数组，转化为有序的树
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
