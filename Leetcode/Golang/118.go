package main

import (
	"fmt"
)

func main() {
	test := 7
	fmt.Println(generate(test))
}

func generate(numRows int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, []int{1})
	if numRows == 1 {
		return ans
	}
	ans = append(ans, []int{1, 1})
	if numRows == 2 {
		return ans
	}
	for i := 2; i < numRows; i++ {
		temp := []int{1}
		// fmt.Println(ans[i-1])
		for j := 1; j < len(ans[i-1]); j++ {
			new := ans[i-1][j] + ans[i-1][j-1]
			// fmt.Println(ans[i-1][j])
			// fmt.Println(new)
			temp = append(temp, new)

		}
		temp = append(temp, 1)
		ans = append(ans, temp)

	}

	return ans

}
