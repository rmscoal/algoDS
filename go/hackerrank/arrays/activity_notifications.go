package arrays

import (
	"fmt"
	"sort"
)

func calculateSortedMedian(arr []int32) float32 {
	middle := (len(arr) - 1) / 2

	if len(arr)%2 == 0 {
		return float32(arr[middle]+arr[middle+1]) / 2
	}

	return float32(arr[middle])
}

func insertValueIntoSortedArray(slc []int32, value int32) []int32 {
	// Find the index to insert the value while maintaining sorted order
	index := sort.Search(len(slc), func(i int) bool {
		return slc[i] >= value
	})

	fmt.Println("Index", index)

	// Insert the value at the determined index
	return append(slc[:index], append([]int32{value}, slc[index:]...)...)
}

func removeValueFromSlice(slc []int32, value int32) []int32 {
	// Find the index of the value to remove
	indexToRemove := -1
	for i, v := range slc {
		if v == value {
			indexToRemove = i
			break
		}
	}

	// Check if the value was found and remove it
	if indexToRemove != -1 {
		// Create a new slice without the element to remove
		return append(slc[:indexToRemove], slc[indexToRemove+1:]...)
	}

	return slc
}

func activityNotifications(expenditure []int32, d int32) int32 {
	var notifyCount int32

	windowSlice := append(make([]int32, 0, d), expenditure[0:d-1]...)

	sort.SliceStable(windowSlice, func(i, j int) bool {
		return windowSlice[i] < windowSlice[j]
	})

	for i := int(d); i < len(expenditure); i++ {
		windowSlice = insertValueIntoSortedArray(windowSlice, expenditure[i-1])
		median := calculateSortedMedian(windowSlice)

		if median*2 <= float32(expenditure[i]) {
			notifyCount++
		}

		windowSlice = removeValueFromSlice(windowSlice, expenditure[i-int(d)])
	}

	return notifyCount
}
