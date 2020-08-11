package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// Initialization
	els := []int{9, 8, 7, 6, 5}

	// Execution
	BubbleSort(els)

	// Validation
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])

}

func TestBubbleSortBestCase(t *testing.T) {
	// Initialization
	els := []int{5, 6, 7, 8, 9}

	// Execution
	BubbleSort(els)

	// Validation
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])

}

func getElements(n int) []int {
	result := []int{}
	for i := n; i > 0; i-- {
		result = append(result, i)
	}
	return result
}

// the goal of benchmark is to test the performance of a function by running mamy iterations and take
// an average of each operation
func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}

}

func BenchmarkSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort10000(b *testing.B) {
	els := getElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}

}

func BenchmarkSort10000(b *testing.B) {
	els := getElements(10000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}
