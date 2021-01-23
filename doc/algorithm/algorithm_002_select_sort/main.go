package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	selectSort(arr)
	fmt.Printf("%v", arr)
}

func selectSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minPos := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minPos] {
				minPos = j
			}
		}
		if minPos != i {
			arr[i], arr[minPos] = arr[minPos], arr[i]
		}
	}
}

//func aSort(arr []int) {
//	n := len(arr)
//	for i := 0; i < n; i++ {
//		for j := i + 1; j < n; j++ {
//			if arr[i] > arr[j] {
//				arr[i], arr[j] = arr[j], arr[i]
//			}
//		}
//	}
//}
