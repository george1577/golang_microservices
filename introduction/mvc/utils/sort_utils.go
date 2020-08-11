package utils

import "sort"

func BubbleSort(elements []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
}

func OptimalSort(elements []int) {
	// after testing with benchmark, we figured that 30000 is kind of a threshold, that being said
	// if there are less than 30K elements, using bubble sort is more efficient, otherwise using Go built-in sort function
	// is more efficient
	if len(elements) < 30000 {
		BubbleSort(elements)
	} else {
		sort.Ints(elements)
	}
}
