package sort

func Quicksort(arr []int) {
	quicksort(arr, 0, len(arr)-1)
}

func quicksort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := arr[end]
	left := start
	right := end - 1
	for left < right {
		for arr[left] <= pivot && left < right {
			left++
		}
		for arr[right] >= pivot && left < right {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	if arr[left] >= pivot {
		arr[left], arr[end] = arr[end], arr[left]
	} else {
		left++
	}
	quicksort(arr, start, left-1)
	quicksort(arr, left+1, end)
}
