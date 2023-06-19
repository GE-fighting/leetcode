package main

import "fmt"

// 1、快排  //5,6,8,1,4,3,9,2,7,10    5,2,3,1,4,8,9,6,7,10
func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivot := arr[0]
	low, high := 1, len(arr)-1

	for low <= high {
		for low <= high && arr[low] >= pivot {
			low++
		}

		for low <= high && arr[high] <= pivot {
			high--
		}

		if low <= high {
			arr[low], arr[high] = arr[high], arr[low]
		}
	}

	arr[0], arr[high] = arr[high], arr[0]

	quickSort(arr[:high])
	quickSort(arr[high+1:])
}

func main() {
	value := []int{3, 2, 1, 5, 6, 4}
	quickSort(value)
	fmt.Println(value)
}
