package tree

func RunTree() {
	var myTree Tree
	myTree.Add(5)
	myTree.Add(25)
	myTree.Add(13)
	myTree.Add(1)
	println("Binary Search Tree:")
	myTree.Traverse()
	println()
}
