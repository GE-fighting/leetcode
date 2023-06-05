package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	record := map[byte]byte{}
	reverseRecord := map[byte]byte{}
	length := len(s)
	for i := 0; i < length; i++ {
		v := record[s[i]]
		if v != 0 {
			if v != t[i] {
				return false
			}
		} else {
			reverseV := reverseRecord[t[i]]
			if reverseV != s[i] {
				return false
			}
			reverseRecord[t[i]] = s[i]
			record[s[i]] = t[i]
		}

	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	fmt.Sprint(isIsomorphic(s, t))
}
