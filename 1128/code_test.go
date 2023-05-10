package main

import (
	"fmt"
	"testing"
)

func TestNumEquivDominoPairs(t *testing.T) {
	dominoes := [][]int{
		[]int{1, 2},
		[]int{1, 2},
		[]int{1, 1},
		[]int{1, 2},
		[]int{2, 2},
	}
	res := numEquivDominoPairs(dominoes)
	fmt.Println(res)
}
