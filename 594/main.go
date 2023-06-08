package main

import "sort"

func findLHS(nums []int) int {
	record := make(map[int]int)
	for _, v := range nums {
		record[v] = record[v] + 1
	}
	keys := []int{}
	for key := range record {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	res := 0
	for i := 1; i < len(keys); i++ {
		if keys[i]-keys[i-1] == 1 {
			temp := record[keys[i]] + record[keys[i-1]]
			if temp > res {
				res = temp
			}
		}
	}
	return res
}
