package main

import (
	"fmt"
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Slice(houses, func(i, j int) bool {
		return houses[i] < houses[j]
	})
	sort.Slice(heaters, func(i, j int) bool {
		return heaters[i] < heaters[j]
	})
	before := []int{}
	in := []int{}
	after := []int{}
	for i := 0; i < len(houses); i++ {
		if houses[i] < heaters[0] {
			before = append(before, houses[i])
		} else if houses[i] > heaters[len(heaters)-1] {
			after = append(after, houses[i])
		} else {
			in = append(in, houses[i])
		}
	}
	low := 0
	beforeValue := 0
	inValue := 0
	afterValue := 0
	if len(before) > 0 {
		beforeValue = heaters[0] - before[0]
	}
	if len(after) > 0 {

		afterValue = after[len(after)-1] - heaters[len(heaters)-1]
	}
	if len(in) > 0 {
		for i := 0; i < len(heaters); i++ {
			if i+1 < len(heaters) {
				tmp := []int{}
				for low < len(in) && in[low] < heaters[i+1] {
					tmp = append(tmp, in[low])
					low++
				}
				if len(tmp) > 0 {
					value := getMinRAadius(heaters[i], heaters[i+1], tmp)
					if value > inValue {
						inValue = value
					}
				}

			}

		}
	}
	res := beforeValue
	if afterValue > res {
		res = afterValue
	}
	if inValue > res {
		res = inValue
	}
	return res
}

func getMinRAadius(low, high int, record []int) int {
	minR := math.MaxInt32

	//i 交于low暖器供暖的个数
	for i := 0; i <= len(record); i++ {
		a := 0
		b := 0
		if i > 0 {
			a = record[i-1] - low
			if i == len(record) {
				b = 0
			} else {
				b = high - record[i]
			}
		} else {
			b = high - record[0]
		}
		value := a
		if b > value {
			value = b
		}
		if value < minR {
			minR = value
		}

	}
	return minR
}

func main() {
	houses := []int{581030105, 557810404, 146319451, 908194298, 500782188, 657821123}
	heater := []int{102246882, 269406752, 816731566, 884936716, 807130337, 578354438}
	fmt.Println(findRadius(houses, heater))
}
