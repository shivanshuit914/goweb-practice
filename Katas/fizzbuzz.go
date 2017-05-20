package main

import "strconv"

func fizzBuzz(input int) string {
	var output string
	if input%3 == 0 {
		output += "fizz"
	}

	if input%5 == 0 {
		output += "buzz"
	}

	if input%3 != 0 && input%5 != 0 {
		output += strconv.Itoa(input)
	}

	return output
}
