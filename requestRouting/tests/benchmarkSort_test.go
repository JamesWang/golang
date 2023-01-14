package tests

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sizes := []int{10, 100, 250}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Array size %v", size), func(subB *testing.B) {
			data := make([]int, size)
			subB.ResetTimer()
			for i := 0; i < subB.N; i++ {
				subB.StopTimer()
				for j := 0; j < size; j++ {
					data[j] = rand.Int()
				}
				subB.StartTimer()
				SortAndTotal(data)
			}
		})
	}
}
