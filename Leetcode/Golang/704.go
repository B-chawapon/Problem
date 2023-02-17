package main

func search(nums []int, target int) int {

	temp := recur(nums, target, 0, len(nums)-1)

	return temp

}

func recur(nums []int, target int, left int, right int) int {
	if right == -1 || left > right {
		return -1
	}

	index_mid := (right + left) / 2
	// fmt.Println("--X", index_mid, nums[index_mid], nums, left, right)
	if target == nums[index_mid] {
		return index_mid
	} else if target > nums[index_mid] {
		return recur(nums, target, index_mid+1, right)

	} else if target < nums[index_mid] {
		return recur(nums, target, left, index_mid-1)
	}
	return -1
}
