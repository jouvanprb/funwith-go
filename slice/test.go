package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Contoh Input nama ke slice data
// func main() {
// 	sliDt := make([]string, 1,3)
// 	for {
// 		fmt.Print("-> ")
// 		reader := bufio.NewReader(os.Stdin)
// 		input,_ := reader.ReadString('\n')
// 		input = strings.TrimSpace(input)

// 		if input == "x" || input == "X" {
// 			fmt.Println("Program Stopped")
// 			break
// 		}

// 		sliDt = append(sliDt, input)
// 	}
// 	fmt.Println(sliDt)
// }


func main() {
	sliDt := make([]int, 0,3)
	for {
		fmt.Print("-> ")
		// ini fungsi untuk melakukan inputan string by terminal
		reader := bufio.NewReader(os.Stdin)
		input,_ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// kondisi untuk men stop loop program apabila menginput sesuai kondisi
		if input == "x" || input == "X" {
			fmt.Println("Program Stopped")
			break
		}

		// meng convert nilai string menjadi integer
		convInt, err := strconv.Atoi(input) 
		// kondisi disiini berfungsi jika ada nilai selain int tidak di masukkan ke slice data -> continue
		if err != nil {
			fmt.Println("Return error:", err)
			continue
		}

		sliDt = append(sliDt, convInt)
		// Ini berfungsi sebagai meng sort tiap data
		sort.Ints(sliDt)
		fmt.Println(sliDt)
	}
}