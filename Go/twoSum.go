package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Printf("[%d,%d]\n", i, j)
			}
		}
	}
}
