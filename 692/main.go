package main

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	wordCount := map[string]int{}
	for _, v := range words {
		wordCount[v] = wordCount[v] + 1
	}
	resUnSort := []string{}
	for k, _ := range wordCount {
		resUnSort = append(resUnSort, k)
	}
	sort.Slice(resUnSort, func(i, j int) bool {
		if wordCount[resUnSort[i]] > wordCount[resUnSort[j]] {
			return true
		} else if wordCount[resUnSort[i]] == wordCount[resUnSort[j]] {
			return dictionarySort(resUnSort[i], resUnSort[j])
		}
		return false
	})
	return resUnSort[0:k]
}

func dictionarySort(a, b string) bool {
	lenA := len(a)
	lenB := len(b)
	length := lenA
	if lenB < lenA {
		length = lenB
	}
	for i := 0; i < length; i++ {
		if a[i] > b[i] {
			return false
		} else if a[i] < b[i] {
			return true
		}
	}
	if lenA > lenB {
		return false
	}
	return true
}

func main() {
	words := []string{"a", "aa", "aaa"}
	frequent := topKFrequent(words, 2)
	fmt.Println(frequent)
	fmt.Println("a" > "aa")

}
