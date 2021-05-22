/*
| Write a program to sort an array of integers. The program should partition the array into 4
| parts, each of which is sorted by a different goroutine. Each partition should be of
| approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into
| one large sorted array.
|
| The program should prompt the user to input a series of integers. Each goroutine which sorts
| ¼ of the array should print the subarray that it will sort. When sorting is complete, the
| main goroutine should print the entire sorted list.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	items, err := getItems()
	if err != nil {
		panic(err)
	}

	ch := make(chan []int)

	q11, q12, q21, q22 := Split4(items)

	go goSort(q11, ch)
	go goSort(q12, ch)
	go goSort(q21, ch)
	go goSort(q22, ch)

	a, b, c, d := <-ch, <-ch, <-ch, <-ch

	sorted := Sort(Merge(a, b, c, d))
	fmt.Println(sorted)
}

func Merge(slices ...[]int) []int {
	merged := make([]int, 0)
	for _, slice := range slices {
		merged = append(merged, slice...)
	}
	return merged
}

func goSort(items []int, c chan []int) {
	fmt.Println(items)
	c <- Sort(items)
}

func Split4(items []int) ([]int, []int, []int, []int) {
	len := len(items)
	if len == 1 {
		return items, []int{}, []int{}, []int{}
	}
	if len == 2 {
		return items[:1], items[1:], []int{}, []int{}
	}
	if len == 3 {
		return items[:1], items[1:2], items[2:3], []int{}
	}

	n2 := len / 2
	n4 := len / 4

	return items[:n4], items[n4:n2], items[n2 : n2+n4], items[n2+n4:]
}

func getItems() ([]int, error) {
	cli := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter up to 10 integers")
	if !cli.Scan() {
		return nil, cli.Err()
	}

	input := strings.Fields(strings.TrimSpace(cli.Text()))
	items := make([]int, len(input))
	var err error
	for i := 0; i < len(input); i++ {
		if items[i], err = strconv.Atoi(input[i]); err != nil {
			return nil, err
		}
	}

	return items, nil
}

func Sort(items []int) []int {
	sorted := true

	for i := 0; i < len(items)-1; i++ {
		if items[i] > items[i+1] {
			sorted = false
			Swap(&items, i)
		}
	}

	if !sorted {
		return Sort(items)
	} else {
		return items
	}
}

func Swap(pointer *[]int, index int) {
	items := *pointer
	item := items[index+1]
	items[index+1] = items[index]
	items[index] = item
}
