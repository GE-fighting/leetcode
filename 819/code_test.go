package main

import (
	"fmt"
	"testing"
)

func TestMostCommonWord(t *testing.T) {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	res := mostCommonWord(paragraph, banned)
	fmt.Println(res)
}
