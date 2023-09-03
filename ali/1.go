package main

import (
	"fmt"
)

var record []int

func f(n int) int {
	if n == 0 {
		return 1
	} else {
		if n%2 == 0 && n != 1 {
			if record[n%2] != 0 {
				record[n] = record[n%2]
			} else {
				fmt.Printf("n is ---%d\n", n)
				record[n] = f(n % 2)
			}
			return record[n]
		}
		if record[n%2] != 0 || n == 1 {
			record[n] = record[n%2] * 2
		} else {
			record[n] = f(n%2) * 2
		}
		return record[n]
	}
}

func main() {
	var n int
	var a int
	fmt.Scan(&n)
	record = make([]int, 1000000000)
	res := map[int]int{}
	for n > 0 {
		fmt.Scan(&a)
		res[f(a)] = 1
		n--
	}
	fmt.Println(len(res))
}
