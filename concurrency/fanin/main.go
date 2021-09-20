package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://github.com/lotusirous/go-concurrency-patterns/blob/main/4-fanin/main.go

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

// merge N channels into 1 channel
func fanIn(cs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ci := range cs {
		go func(cc <-chan string) {
			for v := range cc { // 当cc被close时, 这个goroutine退出
				c <- v
			}
		}(ci)
	}
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ahn"), boring("Bill"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("exit")
}
