package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("--- 01. ---")
	nums := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(nums, target)
	fmt.Printf("[%v]", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), ","), "[]"))
	fmt.Println("\n")
}
