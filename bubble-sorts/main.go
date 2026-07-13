
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numberSlice := make([]int, 0, 10)
	reader := bufio.NewReader(os.Stdin)

	for i := range 10 {
		fmt.Printf("%d. ", i+1)

		numberStr, err := reader.ReadString('\n')
		numberStr = strings.TrimSpace(numberStr)
		if err != nil {
			fmt.Println("Return Erorr: ", err)
		}
		number, _ := strconv.Atoi(numberStr)
		numberSlice = append(numberSlice, number)
		BubbleSort(numberSlice)
	}
	fmt.Println(numberSlice)
}

func BubbleSort(number []int) {
	n := len(number)
	for _ = range n {
		for j := range n - 1 {
			if number[j] > number[j+1] {
				// number[j], number[j+1] = number[j+1], number[j]
				Swap(number, j)
			}
		}
	}
}

func Swap(number []int, j int) {
	number[j], number[j+1] = number[j+1], number[j]
}
