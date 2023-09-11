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

// 2、堆排序
func heapify(arr []int, n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func heapSort(arr []int) {
	n := len(arr)

	// 建立堆（最大堆）
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 逐个将堆顶元素与最后一个元素交换，并重新构建堆
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func main() {
	value := []int{3, 2, 1, 5, 6, 4}
	heapSort(value)
	fmt.Println(value)
}
