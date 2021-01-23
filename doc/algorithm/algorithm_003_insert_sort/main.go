package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	insertSort(arr)
	fmt.Printf("%v", arr)
}

func insertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		pos := i
		for j := i - 1; j >= 0; j-- {
			if arr[i] < arr[j] {
				pos = j
			} else {
				break
			}
		}
		if pos != i {
			tmp := arr[i]
			for k := i; k > pos; k-- {
				arr[k], arr[k-1] = arr[k-1], arr[k]
			}
			arr[pos] = tmp
		}
	}
}
