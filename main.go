package main

import "go-algorithm/linkedlist"

func main() {
	l := linkedlist.NewList()
	//l.Add(2).Add(5).Add(3).PrintAllNode()
	//l.Add(2).Add(5).Add(3).ReverseDirect().PrintAllNode()
	//l.Add(2).Add(5).Add(3).ReverseRecursive().PrintAllNode()
	l.Add(2).Add(5).Add(3).ReverseInsert().PrintAllNode()
}
