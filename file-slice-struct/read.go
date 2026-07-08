package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Make a struct type contain firstName & lastName
	type name struct {
		Fname string
		Lname string
	}

	// Make slice of struct to collect the data
	var people []name

	// User input name of file by terminal
	var file string
	fmt.Print("Enter file name: ")
	fmt.Scan(&file)
	
	// Open file with os.Open
	// This way to get more control, with buffer and os.open
	data, err := os.Open(file)
	if err != nil {
		fmt.Printf("File not found %s : %v\n", file,err)
		return
	}
	// Close file after func complete exec
	defer data.Close()

	// To scan file
	scanner := bufio.NewScanner(data)

	// Read file line by line
	for scanner.Scan() {
		line := scanner.Text()

		// make parts var for separate by space, so every contain in file has 2 parts
		// Format: "FirstName LastName"
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			continue
		}

		// Create name struct from parts
		// parts 0 for first name and parts 1 for last name
		fullName := name{parts[0], parts[1]}
		// add data to people slice
		people = append(people, fullName)
	}
	// Output
	fmt.Println(people)
}
