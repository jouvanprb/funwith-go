package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Using bufio reader for standar input
	reader := bufio.NewReader(os.Stdin)

	// Input name  
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Enter age with string value
	fmt.Print("Enter Age: ")
	ageStr,_ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)

	// Convert age string value to int
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("Return Error: ", err)
		return
	}

	// Input for address
	fmt.Print("Enter Address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	// Make object with struct for collect the data from input
	type user struct {
		Name    string
		Age     int
		Address string
	}

	// Insert the data from input to the object
	dataUser := user{
		Name:    name,
		Age:     age,
		Address: address,
	}

	// Convert the data in object to JSON
	barr, err := json.Marshal(dataUser)
	if err != nil {
		fmt.Println("Return Error:", err)
	}

	// This is the output
	// fmt.Println(barr)
	fmt.Println(string(barr))

}
