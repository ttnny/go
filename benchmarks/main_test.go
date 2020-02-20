package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// Test case 1: Worst
func TestBubbleSortWorstCase(t *testing.T) {
	// Init
	data := []int{9, 8, 7, 6, 5}

	// Execute
	bubbleSort(nil)

	// Validate
	assert.NotNil(t, data)
	assert.EqualValues(t, 5, len(data))
	assert.EqualValues(t, 5, data[0])
	assert.EqualValues(t, 6, data[1])
	assert.EqualValues(t, 7, data[2])
	assert.EqualValues(t, 8, data[3])
	assert.EqualValues(t, 9, data[4])
}

// Test case 2: Best
func TestBubbleSortBestCase(t *testing.T) {
	// Init
	data := []int{5, 6, 7, 8, 9}

	// Execute
	bubbleSort(data)

	// Validate
	assert.NotNil(t, data)
	assert.EqualValues(t, 5, len(data))
	assert.EqualValues(t, 5, data[0])
	assert.EqualValues(t, 6, data[1])
	assert.EqualValues(t, 7, data[2])
	assert.EqualValues(t, 8, data[3])
	assert.EqualValues(t, 9, data[4])
}

// Helper for benchmarking
func getData(n int) []int {
	result := make([]int, n)

	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}

	return result
}

func BenchmarkBubbleSort10(b *testing.B) {
	data := getData(10)

	for i := 0; i < b.N; i++ {
		bubbleSort(data)
	}
}

// Native sort in Go
func BenchmarkSort10(b *testing.B) {
	data := getData(10)

	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	data := getData(1000)

	for i := 0; i < b.N; i++ {
		bubbleSort(data)
	}
}

// Native sort in Go
func BenchmarkSort1000(b *testing.B) {
	data := getData(1000)

	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	data := getData(100000)

	for i := 0; i < b.N; i++ {
		bubbleSort(data)
	}
}

// Native sort in Go
func BenchmarkSort100000(b *testing.B) {
	data := getData(100000)

	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}
