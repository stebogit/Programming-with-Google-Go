/*
| Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence
| of up to 10 integers. The program should print the integers out on one line, in sorted order,
| from least to greatest. Use your favorite search tool to find a description of how the bubble
| sort algorithm works.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
| As part of this program, you should write a function called BubbleSort() which takes a slice
| of integers as an argument and returns nothing. The BubbleSort() function should modify the
| slice so that the elements are in sorted order.
*/
func BubbleSort(pointer *[]int) {
	items := *pointer
	sorted := true

	for i := 0; i < len(items)-1; i++ {
		if items[i] > items[i+1] {
			sorted = false
			Swap(&items, i)
		}
	}

	if !sorted {
		BubbleSort(&items)
	}
}

/*
| A recurring operation in the bubble sort algorithm is the Swap operation which swaps the
| position of two adjacent elements in the slice. You should write a Swap() function which
| performs this operation. Your Swap() function should take two arguments, a slice of integers
| and an index value i which indicates a position in the slice. The Swap() function should return
| nothing, but it should swap the contents of the slice in position i with the contents in
| position i+1.
*/
func Swap(pointer *[]int, index int) {
	items := *pointer
	item := items[index+1]
	items[index+1] = items[index]
	items[index] = item
}

func main() {
	cli := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter up to 10 integers")

	if cli.Scan() {
		input := strings.Fields(strings.TrimSpace(cli.Text()))
		items := make([]int, len(input))
		var err error
		for i := 0; i < len(input); i++ {
			fmt.Println(input[i])
			if items[i], err = strconv.Atoi(input[i]); err != nil {
				panic(err)
			}
		}

		BubbleSort(&items)

		fmt.Println(items)

	} else {
		fmt.Println(cli.Err().Error())
	}
}
