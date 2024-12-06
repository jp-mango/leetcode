package main

import "slices"

func searchInsert(nums []int, target int) (int, bool) {
	n, found := slices.BinarySearch(nums, target)
	return n, found
}
