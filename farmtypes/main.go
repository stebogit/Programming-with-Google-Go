/*
| Write a program which allows the user to create a set of animals and to get information about those animals.
| Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create
| a new animal of one of the three types, or the user can request information about an animal that he/she has
| already created. Each animal has a unique name, defined by the user. Note that the user can define animals of
| a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table
| contains the three types of animals and their associated data.
|
|   | Animal | Locomotion method | Food eaten | Spoken sound |
|   +--------+-------------------+------------+--------------+
|   | cow    | walk              | grass      | moo          |
|   | bird   | fly               | worms      | peep         |
|   | snake  | slither           | mice       | hsss         |
|
| Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your
| program should accept one command at a time from the user, print out a response, and print out a new prompt on
| a new line. Your program should continue in this loop forever. Every command from the user must be either a
| “newanimal” command or a “query” command.
|
| Each “newanimal” command must be a single line containing three strings:
|   - the first string is “newanimal”
|   - the second string is an arbitrary string which will be the name of the new animal
|   - the third string is the type of the new animal, either “cow”, “bird”, or “snake”
| Your program should process each newanimal command by creating the new animal and printing “Created it!” on
| the screen.
|
| Each “query” command must be a single line containing 3 strings.
|   - the first string is “query”
|   - the second string is the name of the animal
|   - the third string is the name of the information requested about the animal (“eat”, “move”, or “speak”)
| Your program should process each query command by printing out the requested data.
|
| Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal
| interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
| The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the
| Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these
| three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the
| Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should
| call the appropriate method when the user issues a query command.
| */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"coursera/farmtypes/animals"
)

var farm = make(map[string]animals.Animal)

func NewAnimal(name, animalType string) {
	switch animalType {
	case "cow":
		farm[name] = animals.Cow{}
	case "bird":
		farm[name] = animals.Bird{}
	case "snake":
		farm[name] = animals.Snake{}
	default:
		fmt.Println("Invalid query entered!")
		return
	}

	fmt.Println("Created it!")
}

func Query(name, action string) {
	switch action {
	case "eat":
		farm[name].Eat()
	case "move":
		farm[name].Move()
	case "speak":
		farm[name].Speak()
	default:
		fmt.Println("Invalid query entered!")
	}
}

func main() {
	cli := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter your command:")
		fmt.Print("> ")

		if cli.Scan() {
			request := strings.Fields(strings.TrimSpace(cli.Text()))
			command := strings.ToLower(request[0])
			name := strings.ToLower(request[1])

			if command == "newanimal" {
				animalType := strings.ToLower(request[2])
				NewAnimal(name, animalType)
			} else if command == "query" {
				action := strings.ToLower(request[2])
				Query(name, action)
			} else {
				fmt.Println("Unrecognized command")
			}
		} else {
			fmt.Println(cli.Err().Error())
		}
	}
}
