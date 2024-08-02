package main

import "fmt"

func main() {
	arr := []int{4, 3, 0, 1, 10, 234, 10, -1, 3, -5}
	fmt.Println(quickSort(arr, 0, len(arr)-1))
}

func quickSort(arr []int, start, end int) []int {
	if end-start+1 <= 1 {
		return arr
	}
	pivot := arr[end]
	left := start

	for i := start; i < end+1; i++ {
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[end], arr[left] = arr[left], arr[end]

	quickSort(arr, start, left-1)
	quickSort(arr, left+1, end)

	return arr
}
