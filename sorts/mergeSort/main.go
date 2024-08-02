package main

import "fmt"

func main() {
	arr := []int{4, 3, 0, 1, 10, 234, 10, -1, 3, -5}
	fmt.Println(mergeSort(arr, 0, len(arr)-1))
}

func mergeSort(arr []int, start, end int) []int {
	if end-start+1 <= 1 {
		return arr
	}
	middle := (start + end) / 2
	mergeSort(arr, start, middle)
	mergeSort(arr, middle+1, end)

	return merge(arr, start, middle, end)
}

func merge(arr []int, start, middle, end int) []int {

	leftHalf := make([]int, middle-start+1)
	copy(leftHalf, arr[start:middle+1])
	rightHalf := make([]int, end-middle)
	copy(rightHalf, arr[middle+1:end+1])
	l, r := 0, 0
	k := start

	for l < len(leftHalf) && r < len(rightHalf) {
		if leftHalf[l] <= rightHalf[r] {
			arr[k] = leftHalf[l]
			l++
		} else {
			arr[k] = rightHalf[r]
			r++
		}
		k++
	}
	for l < len(leftHalf) {
		arr[k] = leftHalf[l]
		l++
		k++
	}
	for r < len(rightHalf) {
		arr[k] = rightHalf[r]
		r++
		k++
	}
	return arr
}
