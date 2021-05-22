/*
| Let us assume the following formula for displacement s as a function of time t,
| acceleration a, initial velocity vo, and initial displacement so.
|
| s = ½ a * t^2 + vo * t + so
|
| Write a program which first prompts the user to enter values for acceleration,
| initial velocity, and initial displacement. Then the program should prompt the user
| to enter a value for time and the program should compute the displacement after the
| entered time.
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
| You will need to define and use a function called GenDisplaceFn() which takes three
| float64 arguments, acceleration a, initial velocity vo, and initial displacement so.
| GenDisplaceFn() should return a function which computes displacement as a function of
| time, assuming the given values acceleration, initial velocity, and initial displacement.
| The function returned by GenDisplaceFn() should take one float64 argument t,
| representing time, and return one float64 argument which is the displacement travelled
| after time t.
|
| For example, let’s say that I want to assume the following values for acceleration,
| initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the
| following statement to call GenDisplaceFn() to generate a function fn which will compute
| displacement as a function of time.
| fn := GenDisplaceFn(10, 2, 1)
|
| Then I can use the following statement to print the displacement after 3 seconds.
| fmt.Println(fn(3))
|
| And I can use the following statement to print the displacement after 5 seconds.
| fmt.Println(fn(5))
*/
func GenDisplaceFn(a, v, s float64) func(t float64) float64 {
	dF := func(t float64) float64 {
		return a*t*t/2 + v*t + s
	}
	return dF
}

func main() {
	cli := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter acceleration, initial velocity, and initial displacement:")

	if cli.Scan() {
		input := strings.Fields(strings.TrimSpace(cli.Text()))
		parameters := make([]float64, len(input))
		var err error
		for i := 0; i < len(input); i++ {
			if parameters[i], err = strconv.ParseFloat(input[i], 64); err != nil {
				panic(err)
			}
		}

		fn := GenDisplaceFn(parameters[0], parameters[1], parameters[2])

		fmt.Println("Please enter time:")
		if cli.Scan() {
			if t, err2 := strconv.ParseFloat(strings.TrimSpace(cli.Text()), 64); err != nil {
				panic(err2)
			} else {
				fmt.Println(fn(t))
			}
		} else {
			fmt.Println(cli.Err().Error())
		}

	} else {
		fmt.Println(cli.Err().Error())
	}
}
