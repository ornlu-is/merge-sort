package main

import (
	"fmt"
	"math"
)

func merge(array []int, leftIndex, midIndex, rightIndex int) {
	n1 := midIndex - leftIndex + 1
	n2 := rightIndex - midIndex

	// create and populate left and right subarrays
	leftSubArray := make([]int, n1+1)
	rightSubArray := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		leftSubArray[i] = array[leftIndex+i]
	}
	for j := 0; j < n2; j++ {
		rightSubArray[j] = array[midIndex+j+1]
	}
	// add sentinel value to both arrays
	leftSubArray[n1] = math.MaxInt64
	rightSubArray[n2] = math.MaxInt64

	i := 0
	j := 0
	// merge the arrays while sorting
	for k := leftIndex; k < rightIndex+1; k++ {
		if leftSubArray[i] <= rightSubArray[j] {
			array[k] = leftSubArray[i]
			i++
		} else {
			array[k] = rightSubArray[j]
			j++
		}
	}
}

func mergeSort(array []int, leftIndex, rightIndex int) {
	if leftIndex < rightIndex {
		midIndex := (leftIndex + rightIndex) / 2
		mergeSort(array, leftIndex, midIndex)
		mergeSort(array, midIndex+1, rightIndex)
		merge(array, leftIndex, midIndex, rightIndex)
	}
}

func main() {
	array := []int{34, 2, 17, 11, 39, 4, 11, 20}
	fmt.Println("Given array:", array)
	mergeSort(array, 0, len(array)-1)
	fmt.Println("Sorted array:", array)
}
