package quicksort

func QuickSort(arr *[]int) {
	quickSort(arr, 0, len(*arr)-1)
}

func swap(arr []int, p1 int, p2 int) []int {
	temp := arr[p1]
	arr[p1] = arr[p2]
	arr[p2] = temp
	return arr
}

func partition(arr []int, low int, high int) int {

	pivot := arr[high]

	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr = swap(arr, i, j)
		}
	}
	arr = swap(arr, i+1, high)
	return (i + 1)
}

func quickSort(arr *[]int, low int, high int) {
	if low < high {
		pi := partition(*arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}
