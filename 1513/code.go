package main

import (
	"fmt"
	"math"
	"strings"
)

func numSub(s string) int {
	strSlice := strings.Split(s, "0")
	var res int64 = 0
	for _, v := range strSlice {
		if v != "" {
			n := len(v)
			res += int64(n * (n + 1) / 2)
		}
	}
	i2 := int64(math.Pow(10, 9)) + int64(7)
	i := res % int64(i2)
	return int(i)
}

func main() {
	s := "111111"
	fmt.Println(numSub(s))
}
