package main

import "slices"

func searchInsert(nums []int, target int) int {
	n, _ := slices.BinarySearch(nums, target)
	return n
}
