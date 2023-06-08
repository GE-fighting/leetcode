package main

import "fmt"

func longestWord(words []string) string {
	wordRecord := map[string]bool{}
	wordsMap := map[string]int{}
	for j := 0; j < len(words); j++ {
		wordsMap[words[j]] = 1
	}
	var reslutStr string
	length := 0
	for i := 0; i < len(words); i++ {
		status := getWordStatus(wordRecord, wordsMap, words[i])
		if status {
			if len(words[i]) > length {
				length = len(words[i])
				reslutStr = words[i]
			} else if len(words[i]) == length {
				if words[i] < reslutStr {
					reslutStr = words[i]
				}
			}
		}
	}
	return reslutStr
}

func getWordStatus(wordRecord map[string]bool, wordsMap map[string]int, word string) bool {
	val, ok := wordRecord[word]
	if ok {
		return val
	} else {
		if wordsMap[word] != 1 {
			return false
		}
		if len(word) == 1 {
			wordRecord[word] = true
			return true
		}
	}
	wordSub := string(word[0 : len(word)-1])
	res := getWordStatus(wordRecord, wordsMap, wordSub)
	wordRecord[word] = res
	return res
}

func main() {
	words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	word := longestWord(words)
	fmt.Println(word)
}
