package binarysearch

func RunSearch() {
	arr := []int{5, 10, 20, 3, 7, 1, 12, 100, 55, 6, 8, 9}
	println("Binrary Search for ", 6, " in list:")
	val := Find(arr, 6)
	println(val)
}
