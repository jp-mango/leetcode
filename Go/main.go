package main

import (
	"fmt"
)

func main() {
	fmt.Println("- 01. Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.")
	nums := []int{2, 7, 11, 15}
	fmt.Println("Nums:", nums)
	target := 9
	fmt.Println("Target:", target)
	result := TwoSum(nums, target)
	fmt.Printf("Result: %v\n", result)

	fmt.Println("\n- 02. You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.")

	fmt.Println("\n- 58. Given a string s consisting of words and spaces, return the length of the last word in the string.")
	phrase :=
		`Is i was an ant crawlin' upon the wall Tell me, baby, would it make no difference at all            `
	lastWord, l := lengthOfLastWord(phrase)
	fmt.Printf("Phrase: %s\nLast word: %s\nLength of %s: %d", phrase, lastWord, lastWord, l)
}
