package main

import "fmt"

func main() {
	arr := []int{4, 3, 0, 1, 10, 234, 10, -1, 3, -5}
	fmt.Println(sort(arr))
}

func sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for j >= 0 && arr[j+1] < arr[j] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			j--
		}
	}
	return arr
}
