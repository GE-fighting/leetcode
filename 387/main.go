package main

import (
	"fmt"
	"math"
)

func firstUniqChar(s string) int {
	record := map[byte]int{}
	recordIndex := map[byte]int{}
	for i := 0; i < len(s); i++ {
		record[s[i]] = record[s[i]] + 1
		recordIndex[s[i]] = i
	}
	min := math.MaxInt32
	for k, v := range record {
		if v == 1 {
			if recordIndex[k] < min {
				min = recordIndex[k]
			}
		}
	}
	if min == math.MaxInt32 {
		return -1
	}
	return min
}

func main() {
	s := "leetcode"
	char := firstUniqChar(s)
	fmt.Println(char)
}
