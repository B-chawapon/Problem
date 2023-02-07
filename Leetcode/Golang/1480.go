package main

func runningSum(nums []int) []int {
	temp := nums[0]
	for i := 1; i < len(nums); i++ {
		temp += nums[i]
		nums[i] = temp
	}
	return nums

}
