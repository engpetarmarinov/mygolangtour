package sorting

func QuickSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}

	low := 0
	high := len(arr) - 1
	quickSortInPlace(arr, low, high)
}

func quickSortInPlace(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSortInPlace(arr, low, pivotIndex-1)
		quickSortInPlace(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func QuickSortAlloc(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}
	equal := []int{}

	for _, value := range arr {
		switch {
		case value < pivot:
			left = append(left, value)
		case value > pivot:
			right = append(right, value)
		default:
			equal = append(equal, value)
		}
	}

	left = QuickSortAlloc(left)
	right = QuickSortAlloc(right)

	left = append(left, equal...)
	left = append(left, right...)

	return left
}
