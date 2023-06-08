package main

func uniqueOccurrences(arr []int) bool {
	record := map[int]int{}
	for _, v := range arr {
		record[v] = record[v] + 1
	}
	countMap := map[int]bool{}
	for _, v := range record {
		if _, ok := countMap[v]; !ok {
			countMap[v] = true
		} else {
			return false
		}
	}
	return true
}
