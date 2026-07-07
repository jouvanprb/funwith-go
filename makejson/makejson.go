package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// For accept input with spaces
	reader := bufio.NewReader(os.Stdin)

	// for input string
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// for input addres
	fmt.Print("Enter your addres: ")
	addr, _ := reader.ReadString('\n')
	addr = strings.TrimSpace(addr)

	// For collect key value object
	dataMap := make(map[string]string)

	dataMap["name"] = name
	dataMap["address"] = addr

	// To access and input the data to the map object
	// dataMap[name] = addr
	// Debug the output
	fmt.Println(dataMap)

	// convert the data in object to the JSON
	barr, err := json.Marshal(dataMap)

	// Error handling
	if err != nil {
		panic(err)
	}

	// Output result
	fmt.Println(string(barr))
}
