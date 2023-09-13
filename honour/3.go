package main

import (
	"fmt"
)

func isMatch(a, b string) int {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		}
		if a[i] != b[j] {
			if b[j] == '*' {
				if i == 0 {
					return 0
				}
				if a[i] != a[i-1] {
					return 0
				}
				i++
			} else if b[j] == '.' {
				i++
				j++
			} else {
				return 0
			}

		}
	}

	return 1
}

func main() {
	var strA string
	var strB string
	fmt.Scan(&strA)
	fmt.Scan(&strB)
	fmt.Println(isMatch(strA, strB))
}
