package main

import (
	"fmt"
)

func Main() {
	data2 := []int{1, 2, 5}
	test := 10
	fmt.Println(coinChange(data2, test))
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	return 1

}
