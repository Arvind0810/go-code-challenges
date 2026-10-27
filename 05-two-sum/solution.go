package twosum

import "slices"

// TwoSum returns the indices [i, j], with i < j, of the two elements of nums
// that add up to target. It returns nil if no such pair exists.
func TwoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j >= 0; j-- {
			if slices.Contains(result, i) || slices.Contains(result, j) || i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
			}
		}
	}
	return result
}
