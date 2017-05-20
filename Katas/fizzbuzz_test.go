package main

import "testing"

func testFizzBuzz(test *testing.T) {
	dataProvider := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, "fizz"},
		{5, "buzz"},
		{9, "fizz"},
		{15, "fizzbuzz"},
		{100, "buzz"},
		{3000, "fizzbuzz"},
	}

	for _, data := range dataProvider {
		output := fizzBuzz(data.input)
		if output != data.expected {
			test.Errorf("Test fail Expected: %d Actual: %d", data.expected, output)
		}
	}

}
