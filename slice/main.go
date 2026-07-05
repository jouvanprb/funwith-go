package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// Make slice with cap 3 and len 0
	sli := make([]int, 0, 3)

	fmt.Println("Enter an integer, press 'x' to stop")

	for {
		// variable for scan every input
		var input string

		fmt.Print(">: ")
		fmt.Scan(&input)

		if input == "x" || input == "X" {
			fmt.Println("Program stopped")
			break
		}

		// multiple return, conv to integer
		result, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Return Error :", err)
			continue
		}

		// append the result to sli
		sli = append(sli, result)
		// sort function to sorts integers in slice data 
		sort.Ints(sli)
		// Output sequence
		fmt.Printf("%d\n", sli)
		// fmt.Printf("Result :%d\n" ,result)
	}
}
