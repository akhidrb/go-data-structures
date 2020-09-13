package main

import (
	"playground/data-structures/mylib/linkedlist"
	"playground/data-structures/mylib/tree"
)

func main() {

	var myLink linkedlist.LinkedList
	myLink.Add(5)
	myLink.Add(25)
	myLink.Add(13)
	myLink.Add(1)

	println("LinkedList:")
	pointer := myLink.Head
	for pointer != nil {
		print(pointer.Val, " ")
		pointer = pointer.Next
	}
	println()

	println("Binary Search Tree:")
	var myTree tree.Tree
	myTree.Add(5)
	myTree.Add(25)
	myTree.Add(13)
	myTree.Add(1)
	myTree.Traverse()
}
