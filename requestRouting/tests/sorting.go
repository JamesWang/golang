package tests

import (
	"fmt"
	"sort"
)

func SortAndTotal(vals []int) (sorted []int, total int) {
	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)
	for _, val := range sorted {
		total += val
		//total++
	}
	return
}

func RunSortAndTotal() {
	nums := []int{100, 20, 1, 7, 84}
	sorted, total := SortAndTotal(nums)
	fmt.Println("Sorted Data:", sorted)
	fmt.Println("Total:", total)
}
