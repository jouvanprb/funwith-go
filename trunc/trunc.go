package main

import "fmt"

func main() {
	// declare the var for the floating point
	var floatNum float32

	// input by user in terminal
	fmt.Print("Input :")
	fmt.Scan(&floatNum)
	
	// return the output in integers
	output := int(floatNum)
	fmt.Println(output)
}

