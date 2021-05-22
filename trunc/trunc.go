/*
| Write a program which prompts the user to enter a floating point number and prints the integer which
| is a truncated version of the floating point number that was entered. Truncation is the process of
| removing the digits to the right of the decimal place.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Println("Please enter a float number")
	if _, err := fmt.Scan(&input); err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		if num, err := strconv.ParseFloat(input, 64); err != nil {
			fmt.Printf("Error: %s", err)
		} else {
			fmt.Println(int(num))
		}
	}
}
