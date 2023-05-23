package main

import "strings"

func mostCommonWord(paragraph string, banned []string) string {
	markStr := "!?',;."
	str := []byte(paragraph)
	banned = append(banned, "")
	for i := 0; i < len(str); i++ {
		if strings.Contains(markStr, string(str[i])) {
			str[i] = ' '
		}
	}
	bannedMap := map[string]int{}
	wordMap := map[string]int{}
	for _, v := range banned {
		bannedMap[v] = 1
	}
	words := strings.Split(string(str), " ")
	maxCount := 0
	maxWord := ""
	for _, v := range words {
		v = strings.ToLower(v)
		if bannedMap[v] == 0 {
			wordMap[v] = wordMap[v] + 1
			if wordMap[v] > maxCount {
				maxCount = wordMap[v]
				maxWord = v
			}
		}

	}
	return maxWord
}
