package main

import "fmt"

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	go say("hello")
	say("hello world")
}
