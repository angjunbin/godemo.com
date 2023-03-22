package main

import (
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
	b.ResetTimer()
	data := []int{1, 0, 2, 8, 2, 0, 1, 8}
	for i := 0; i < b.N; i++ {
		QuickSort(data)
	}
}
