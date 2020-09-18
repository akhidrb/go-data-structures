package linkedlist

func RunLinkedList() {
	var myLink LinkedList
	myLink.Add(5)
	myLink.Add(25)
	myLink.Add(13)
	myLink.Add(1)
	println("LinkedList:")
	myLink.Traverse()
	println()
}
