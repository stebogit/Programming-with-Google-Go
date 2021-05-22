/*
| Write a program which prompts the user to enter a string. The program searches through the entered string for
| the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the
| character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!”
| otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
|
| Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”,
| “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func found(str string) bool {
	target := strings.ToLower(str)
	return strings.HasPrefix(target, "i") &&
		strings.HasSuffix(target, "n") &&
		strings.ContainsAny(target, "a")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter a string")

	if scanner.Scan() {
		input := scanner.Text()
		if found(input) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	} else {
		fmt.Println(scanner.Err().Error())
	}
}
