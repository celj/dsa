package binary_search

func Run() any {
	arr := []int{11, 12, 22, 25, 34, 64, 90}
	target := 25
	index := SearchInt(arr, target)
	return map[string]any{
		"array":  arr,
		"target": target,
		"index":  index,
	}
}

func Search[T comparable](arr []T, target T, less func(T, T) bool) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if less(arr[mid], target) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func SearchInt(arr []int, target int) int {
	return Search(arr, target, func(a, b int) bool { return a < b })
}

func SearchString(arr []string, target string) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func SearchFloat64(arr []float64, target float64) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func SearchGeneric[T comparable](arr []T, target T, less func(T, T) bool) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if less(arr[mid], target) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
