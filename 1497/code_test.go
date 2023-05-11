package main

import "testing"

func TestCanArrange(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
	k := 5
	canArrange(arr, k)
}
