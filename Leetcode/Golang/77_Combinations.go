package main

func combine(n int, k int) [][]int {
	answer := make([][]int, 0)

	var listComb func(startlist []int, start int)

	listComb = func(startlist []int, start int) {

		if len(startlist) == k {
			dst := make([]int, k)
			copy(dst, startlist)
			answer = append(answer, dst)
			return
		}

		for i := start; i <= n; i++ {
			startlist = append(startlist, i)
			listComb(startlist, i+1)
			startlist = startlist[:len(startlist)-1]
		}
		return

	}
	listComb(make([]int, 0), 1)

	return answer
}
