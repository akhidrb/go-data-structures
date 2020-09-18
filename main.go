package main

import (
	"playground/data-structures/mylib/binarysearch"
	"playground/data-structures/mylib/linkedlist"
	"playground/data-structures/mylib/queue"
	"playground/data-structures/mylib/quicksort"
	"playground/data-structures/mylib/tree"
)

func main() {

	linkedlist.RunLinkedList()
	queue.RunQueue()
	tree.RunTree()
	binarysearch.RunSearch()
	quicksort.RunQuickSort()
	////////////////////////
	// Run QuickSort with BinarySearch
	arr := []int{5, 10, 20, 3, 7, 1, 12, 100, 55, 6, 8, 9}
	quicksort.QuickSort(&arr)
	println("Sorted Array:")
	for _, val := range arr {
		print(val, " ")
	}
	println()
	pos := binarysearch.Find(arr, 6)
	println("Position of 6 is ", pos)
}
