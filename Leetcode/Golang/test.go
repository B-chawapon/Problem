package main

import "fmt"

func Main() {
	employees := []string{"John", "Mary", "Bob"}

	// for i, employee := range employees {
	// 	fmt.Printf("Employee %d: %s\n", i, employee)
	// }
	for _, employee := range employees {
		fmt.Println(employee)
	}
}
