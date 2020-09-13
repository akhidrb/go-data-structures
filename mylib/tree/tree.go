package tree

type Node struct {
	Val   int
	Right *Node
	Left  *Node
}

type Tree struct {
	Pointer *Node
}

func (t *Tree) Add(val int) {
	if t.Pointer == nil {
		t.Pointer = &Node{
			Val: val,
		}
		return
	}
	t.Pointer.add(val)
}

func (t *Node) add(val int) {
	if val < t.Val {
		if t.Left == nil {
			t.Left = &Node{
				Val: val,
			}
			return
		}
		t.Left.add(val)
	}
	if val > t.Val {
		if t.Right == nil {
			t.Right = &Node{
				Val: val,
			}
			return
		}
		t.Right.add(val)
	}
}

func (t *Tree) Traverse() {
	t.Pointer.traverse()
}

func (n *Node) traverse() {
	if n.Left != nil {
		n.Left.traverse()
	}
	print(n.Val, " ")
	if n.Right != nil {
		n.Right.traverse()
	}
}
