package btree

import (
	"fmt"
	"testing"
)

func TestArr2Tree(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	node := Arr2Tree(arr, 0, len(arr) - 1)
	TreePrintln(node)
}

func TestLevelPrint(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	root := Arr2Tree(arr, 0, len(arr) - 1)
	h := 3
	LevelPrint(root, h)
}

func TestFindMaxSubTree(t *testing.T) {
	arr := []int{1,-2,-3,4,5,-6,7,-8,9,-10}
	node := Arr2Tree(arr, 0, len(arr) - 1)
	FindMaxSubTree(node)
	fmt.Println("max sum", maxSum)
	fmt.Println("max sum node", maxSumNode.Data)
}

func TestIsEqual(t *testing.T) {
	arr1 := []int{1,-2,-3,4,5,-6,7,-8,9,-10}
	arr2 := []int{1,-2,-3,4,5,-6,7,-8,9,-10}
	arr3 := []int{1,-2,-3,4,5,-6,7,-8,9,10}

	root1 := Arr2Tree(arr1, 0, len(arr1) - 1)
	root2 := Arr2Tree(arr2, 0, len(arr2) - 1)
	root3 := Arr2Tree(arr3, 0, len(arr3) - 1)

	var bol bool

	bol = IsEqual(root1, root2)
	if bol == false {
		t.Error("actually is true")
	}

	bol = IsEqual(root1, root3)
	if bol == true {
		t.Error("actually is false")
	}

	bol = IsEqual(root2, root3)
	if bol == true {
		t.Error("actually is false")
	}
}

func TestConvert2linkedList(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	root := Arr2Tree(arr, 0, len(arr) - 1)
	Convert2linkedList(root)
	fmt.Println("开始正向遍历：")
	cur := pHead
	for cur != nil {
		fmt.Print(cur.Data, "->")
		cur = cur.RightChild
	}
	fmt.Println("\n开始反向遍历：")
	cur = pEnd
	for cur != nil {
		fmt.Print(cur.Data, "->")
		cur = cur.LeftChild
	}
}

func TestIsAfterOrder(t *testing.T) {

	arr := [] int{1,3,2,5,8,6,4,7}
	orderMatch := IsAfterOrder(arr, 0, len(arr) - 1)
	if orderMatch == true {
		t.Error("It's real false")
	}

	arr2 := [] int{1,3,2,5,7,6,4}
	orderMatch = IsAfterOrder(arr2, 0, len(arr2) - 1)
	if orderMatch == false {
		t.Error("It's real true")
	}
}