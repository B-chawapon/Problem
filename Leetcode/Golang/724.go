package main

func pivotIndex(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	leftsum := 0
	for i := 0; i < len(nums); i++ {
		if leftsum == sum-leftsum-nums[i] {
			return i
		}
		leftsum += nums[i]
	}
	return -1
}
