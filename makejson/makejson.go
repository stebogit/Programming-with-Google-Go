/*
| Write a program which prompts the user to first enter a name, and then enter an address.
| Your program should create a map and add the name and address to the map using the keys “name”
| and “address”, respectively. Your program should use Marshal() to create a JSON object from
| the map, and then your program should print the JSON object.
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	person := make(map[string]string)

	fmt.Print("Name: ")
	if scanner.Scan() {
		person["name"] = scanner.Text()
	} else {
		fmt.Println(scanner.Err().Error())
		return
	}

	fmt.Print("Address: ")
	if scanner.Scan() {
		person["address"] = scanner.Text()
	} else {
		fmt.Println(scanner.Err().Error())
		return
	}

	if jsonBarr, err := json.Marshal(person); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(string(jsonBarr))
	}
}
