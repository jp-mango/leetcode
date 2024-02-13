package main

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j} // Return a slice of integers directly
			}
		}
	}
	return nil // Return nil if no solution is found
}
