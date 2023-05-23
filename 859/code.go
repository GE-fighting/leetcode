package main

import "fmt"

func buddyStrings(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if a == b {
		byteMap := map[byte]int{}
		for i := 0; i < len(a); i++ {
			byteMap[a[i]] = byteMap[a[i]] + 1
			if byteMap[a[i]] >= 2 {
				return true
			}
		}
		return false
	} else {
		low := 0
		high := len(a) - 1
		for low < high {
			if a[low] != b[low] && a[high] != b[high] {
				break
			}
			if a[low] == b[low] {
				low++
			}
			if a[high] == b[high] {
				high--
			}
		}
		aSlice := []byte(a)
		temp := aSlice[low]
		aSlice[low] = aSlice[high]
		aSlice[high] = temp
		if string(aSlice) == b {
			return true
		}
		return false
	}
}
func main() {
	a := "abab"
	b := "abab"
	flag := buddyStrings(a, b)
	fmt.Println(flag)
}
