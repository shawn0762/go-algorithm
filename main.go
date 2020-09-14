package main

import "fmt"

func main() {
	//a := &[]byte{1, 2, 3}
	//fmt.Printf("%p", a)
	test()
}

func test() {
	a := [3]byte{1,2,3}
	b := make([]byte, 3)

	fmt.Printf("%T", a)
	fmt.Printf("%T", b)
}
