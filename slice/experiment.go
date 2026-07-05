package main

import (
	"fmt"
)

func oke() {
	// sli := make([]int, 3)
	for {
		var input string
		fmt.Print("Input: ")
		fmt.Scan(&input)

		if input == "x" {
			break
		}
		fmt.Println("your input :", input)
	}

}
