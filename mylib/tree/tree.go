package tree

type node struct {
	val   int
	right *node
	left  *node
}

type Tree struct {
	pointer *node
}

func (t *Tree) Add(val int) {
	if t.pointer == nil {
		t.pointer = &node{
			val: val,
		}
		return
	}
	t.pointer.add(val)
}

func (t *node) add(val int) {
	if val < t.val {
		if t.left == nil {
			t.left = &node{
				val: val,
			}
			return
		}
		t.left.add(val)
	}
	if val > t.val {
		if t.right == nil {
			t.right = &node{
				val: val,
			}
			return
		}
		t.right.add(val)
	}
}

func (t *Tree) Traverse() {
	t.pointer.traverse()
}

func (n *node) traverse() {
	if n.left != nil {
		n.left.traverse()
	}
	print(n.val, " ")
	if n.right != nil {
		n.right.traverse()
	}
}
