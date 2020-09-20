package btree

import (
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
