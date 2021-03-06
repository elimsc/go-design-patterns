package main

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	leftmost := make(chan int)
	left, right := leftmost, leftmost
	for i := 0; i < 1000; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Println(<-leftmost)
}
