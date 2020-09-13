package main

import (
	"playground/data-structures/mylib/queue"
	"playground/data-structures/mylib/tree"
)

func main() {

	var myLink queue.Queue
	myLink.Push(5)
	myLink.Push(25)
	myLink.Push(13)
	myLink.Push(1)
	println("Queue:")
	myLink.Traverse()

	println()

	println(myLink.Pop())
	println(myLink.Pop())
	println(myLink.Pop())
	println(myLink.Pop())
	println(myLink.Pop())

	println()

	var myTree tree.Tree
	myTree.Add(5)
	myTree.Add(25)
	myTree.Add(13)
	myTree.Add(1)
	println("Binary Search Tree:")
	myTree.Traverse()
}
