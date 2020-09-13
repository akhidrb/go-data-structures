package stack

type node struct {
	val  int
	next *node
}

type Stack struct {
	head *node
	tail *node
}

func (l *Stack) Push(val int) {
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

func (l *Stack) Pop() int {
	if l.head == nil {
		return 0
	}
	if l.head == l.tail {
		val := l.head.val
		l.head = nil
		return val
	}
	val := l.tail.val
	pointer := l.head
	for pointer.next != l.tail {
		pointer = pointer.next
	}
	l.tail = pointer
	return val
}

func (l *Stack) Traverse() {
	pointer := l.head
	for pointer != nil {
		print(pointer.val, " ")
		pointer = pointer.next
	}
}
