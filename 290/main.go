package main

import "strings"

func wordPattern(pattern string, s string) bool {
	split := strings.Split(s, " ")
	record := map[byte]string{}
	for i := 0; i < len(pattern); i++ {
		if v := record[pattern[i]]; v != "" {
			if v != split[i] {
				return false
			}
		} else {
			record[pattern[i]] = split[i]
		}
	}
	return true
}
