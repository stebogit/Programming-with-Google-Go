/*
| Write a program which reads information from a file and represents it in a slice of structs.
| Assume that there is a text file which contains a series of names. Each line of the text file
| has a first name and a last name, in that order, separated by a single space on the line.
|
| Your program will define a name struct which has two fields, fname for the first name, and lname
| for the last name. Each field will be a string of size 20 (characters).
|
| Your program should prompt the user for the name of the text file. Your program will successively
| read each line of the text file and create a struct which contains the first and last names
| found in the file. Each struct created will be added to a slice, and after all lines have been
| read from the file, your program will have a slice containing one struct for each line in the
| file. After reading all lines from the file, your program should iterate through your slice of
| structs and print the first and last names found in each struct.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	cli := bufio.NewScanner(os.Stdin)
	var path string
	names := make([]Name, 0)

	fmt.Print("File: ")
	if cli.Scan() {
		path = cli.Text()
	} else {
		fmt.Println(cli.Err().Error())
		return
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		items := strings.Split(reader.Text(), " ")
		names = append(names, Name{
			fname: items[0],
			lname: items[1],
		})

	}

	if err := reader.Err(); err != nil {
		fmt.Println(err)
		return
	}

	for _, fullname := range names {
		fmt.Printf("%s %s\n", fullname.fname, fullname.lname)
	}
}
