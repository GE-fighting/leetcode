package main

import (
	"fmt"
	"testing"
)

func TestAddToArrayForm(t *testing.T) {
	num := []int{2, 1, 5}
	k := 806
	result := addToArrayForm(num, k)
	fmt.Println(result)
}
