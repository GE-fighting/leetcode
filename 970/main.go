package main

import (
	"fmt"
	"math"
)

func powerfulIntegers(x int, y int, bound int) []int {
	resMap := map[float64]int{}
	res := []int{}
	x1 := float64(x)
	y1 := float64(y)
	n := 0
	for math.Pow(y1, float64(n)) <= float64(bound) {
		m := 0
		for math.Pow(x1, float64(m))+math.Pow(y1, float64(n)) <= float64(bound) {
			resMap[math.Pow(x1, float64(m))+math.Pow(y1, float64(n))] = 1
			if x == 1 {
				break
			}
			m++
		}
		if y == 1 {
			break
		}
		n++
	}
	for k, _ := range resMap {
		res = append(res, int(k))
	}
	return res
}

func main() {
	x, y, bound := 2, 1, 10
	integers := powerfulIntegers(x, y, bound)
	fmt.Println(integers)
}
