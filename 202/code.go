package main

import "fmt"

func isHappy(n int) bool {
	record := map[int]int{}
	for n != 1 {
		if v := record[n]; v != 0 {
			return false
		}
		record[n] = 1
		n = cal(n)
	}
	return true
}

func cal(n int) int {
	res := 0
	for n != 0 {
		val := n % 10
		res += val * val
		n = n / 10
	}
	return res
}

func main() {
	n := 19
	happy := isHappy(n)
	fmt.Println(happy)
}
