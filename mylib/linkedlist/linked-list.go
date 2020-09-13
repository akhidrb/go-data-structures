package linkedlist

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head *node
	tail *node
}

func (l *LinkedList) Add(val int) {
	if l.head == nil {
		node := node{
			val:  val,
			next: nil,
		}
		l.head = &node
		l.tail = l.head
		return
	}

	node := node{
		val:  val,
		next: nil,
	}
	l.tail.next = &node
	l.tail = l.tail.next
}

func (l *LinkedList) Traverse() {
	pointer := l.head
	for pointer != nil {
		print(pointer.val, " ")
		pointer = pointer.next
	}
}
