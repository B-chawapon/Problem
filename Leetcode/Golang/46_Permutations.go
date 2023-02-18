func permute(nums []int) [][]int {
	answer := make([][]int, 0)

	var listPermute func(startlist []int, start int)

	listPermute = func(startlist []int, start int) {

		if len(startlist) == len(nums) {
			dst := make([]int, k)
			copy(dst, startlist)
			answer = append(answer, dst)
			return
		}

		for i := 0; i < len(nums); i++ {
			startlist = append(startlist, nums[i])
			listPermute(startlist, i+1)
			startlist = startlist[:len(startlist)-1]
		}
		return
	}
	listPermute(nums, 0)

	return answer

}