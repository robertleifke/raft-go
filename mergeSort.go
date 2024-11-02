package main

import (
	"fmt"
)

// Merge function to merge two sorted halves into a sorted array
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Compare elements of left and right slices and append the smaller one
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append the remaining elements of left (if any)
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Append the remaining elements of right (if any)
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

// MergeSort function to sort an array
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Split the array into two halves
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Merge the two sorted halves
	return merge(left, right)
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original array:", arr)

	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
