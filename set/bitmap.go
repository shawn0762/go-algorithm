package set

import (
	"fmt"
	"math"
)

type BitMap struct {
	bits *[]byte
}

// 添加一个元素
func (bs *BitMap) Add(num int) {
	byteLen := len(*bs.bits)
	// 计算落在哪一个字节
	bytePos := num / byteLen
	// 计算落在哪个位
	bitPos := num % 8
	a := *bs.bits
	a[bytePos] |= 1 << bitPos
}

// 判断给定数字是否已存在
func (bs *BitMap) Has(num int) bool {
	bol := false
	byteLen := len(*bs.bits)
	// 计算落在哪一个字节
	bytePos := num / byteLen
	// 计算落在哪个位
	bitPos := num % 8

	bits := *bs.bits
	if bits[bytePos] == (bits[bytePos] | 1 << bitPos) {
		bol = true
	}
	return bol
}

func (bs *BitMap) PrintMap() {
	fmt.Printf("%b", *bs.bits)
}

func NewBitMap(cap int32) *BitMap {
	// 一个byte可以放8个数，
	bytesLen := int32(math.Ceil(float64(cap / 8)))
	b := make([]byte, bytesLen)
	return &BitMap{bits: &b}
}
