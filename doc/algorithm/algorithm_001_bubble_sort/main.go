package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	bubble(arr)
	fmt.Printf("%v", arr)
}

func bubble(arr []int) {
	if arr == nil {
		return
	}
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

func bubble2(arr []int) {
	if arr == nil {
		return
	}
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
