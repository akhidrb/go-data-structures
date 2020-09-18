package quicksort

func RunQuickSort() {
	arr := []int{10, 80, 30, 90, 40, 50, 70}
	QuickSort(&arr)
	println("QuickSort:")
	for _, val := range arr {
		print(val, " ")
	}
	println()
}
