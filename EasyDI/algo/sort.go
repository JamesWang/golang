package sort

func Insertion_Sort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		j := i - 1
		current := arr[i]
		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = current
	}
}

func Bubble_Sort(arr []int) {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
				sorted = false
			}
		}
	}
}
