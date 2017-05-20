package main

import "strconv"

const (
	Fizz       = "fizz"
	FizzNumber = 3
	Buzz       = "buzz"
	BuzzNumber = 5
)

func fizzBuzz(input int) string {
	var output string
	if input%FizzNumber == 0 {
		output += Fizz
	}

	if input%BuzzNumber == 0 {
		output += Buzz
	}

	if input%3 != 0 && input%5 != 0 {
		output += strconv.Itoa(input)
	}

	return output
}
