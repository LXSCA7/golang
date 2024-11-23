package main

import "fmt"

func mySort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func main() {
	list := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(list)
	mySort(list)
	fmt.Println(list)
}
