package main

import "strings"

func lengthOfLastWord(s string) (string, int) {
	wordList := strings.Fields(s)
	lastWord := wordList[len(wordList)-1]
	return lastWord, len(lastWord)
}
