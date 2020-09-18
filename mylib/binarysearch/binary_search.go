package binarysearch

func Find(arr []int, num int) int {
	return find(arr, 0, len(arr)-1, num)
}

func find(arr []int, l int, r int, num int) int {
	if r >= l {
		mid := l + ((r - l) / 2)

		if arr[mid] == num {
			return mid
		}
		if arr[mid] > num {
			return find(arr, l, mid-1, num)
		}
		return find(arr, mid+1, r, num)
	}
	return -1
}
