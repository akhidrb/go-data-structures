package queue

func RunQueue() {
	var myLink Queue
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

}
