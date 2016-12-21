package main

func LongestIncreasingSub(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	longest := 0
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {

		max := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = max + 1

		if dp[i] > longest {
			longest = dp[i]
		}
	}

	return longest
}