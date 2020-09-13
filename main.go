package main

import "playground/data-structures/mylib"

func main() {

	var myLink mylib.LinkedList
	myLink.Add(5)
	myLink.Add(25)
	myLink.Add(13)
	myLink.Add(1)

	pointer := myLink.Head
	for pointer != nil {
		print(pointer.Val, " ")
		pointer = pointer.Next
	}

}
