package set

import (
	"fmt"
)

type BitMap32 struct {
	bits *int32
}

// 添加一个元素
func (bs *BitMap32) Add(num int32) {
	a := *bs.bits
	a = a | (1 << num)
	bs.bits = &a
}

// 判断给定数字是否已存在
func (bs *BitMap32) Has(num int32) bool {
	bol := false
	cur := *bs.bits
	if cur == cur | (int32(1) << num) {
		bol = true
	}
	return bol
}

func (bs *BitMap32) PrintMap() {
	//buffer := &bytes.Buffer{}

	fmt.Printf("%b", *bs.bits)

	//var str []byte
	//
	//for i := 0; i <= 31; i++ {
	//	//var tmp int32
	//	b := *bs.bits
	//	tmp := b & (1 << i)
	//
	//	if tmp == b {
	//		//buffer.WriteString("1")
	//		append(str, "1")
	//		//fmt.Printf("(%d)%d ", i, 1)
	//	} else {
	//		//buffer.WriteString("0")
	//		//fmt.Printf("(%d)%d ", i, 0)
	//	}
	//}
	//fmt.Println(buffer)
}

func NewBitMap32() *BitMap32 {
	i := int32(0)
	return &BitMap32{bits: &i}
}
