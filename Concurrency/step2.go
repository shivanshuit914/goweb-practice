package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

  bufferch := make(chan int, 2)
	bufferch <- 1
	bufferch <- 2
  c3 := func() { bufferch <- 3 }
	go c3()
	fmt.Println(<-bufferch)
	fmt.Println(<-bufferch)
	fmt.Println(<-bufferch)
}
