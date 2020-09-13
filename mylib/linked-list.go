package mylib

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	n    int
}

func (l *LinkedList) Add(val int) {
	if l.n == 0 {
		node := Node{
			Val:  val,
			Next: nil,
		}
		l.Head = &node
		l.Tail = l.Head
		l.n++
		return
	}

	node := Node{
		Val:  val,
		Next: nil,
	}
	l.Tail.Next = &node
	l.Tail = l.Tail.Next
}
