package main

import "fmt"

func isPerfectSquare(num int) bool {
	n := getSquareNum(num)
	if num == n*n {
		return true
	}
	return false
}

func getSquareNum(num int) int {
	low := 1
	high := num
	for high > low+1 {
		if ((low+high)/2)*((low+high)/2) == num {
			return (low + high) / 2
		} else if ((low+high)/2)*((low+high)/2) > num {
			high = (low + high) / 2
		} else {
			low = (low + high) / 2
		}
	}
	return low
}

func main() {
	fmt.Println(isPerfectSquare(104976))
}
