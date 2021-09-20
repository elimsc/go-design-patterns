package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit: // 调用者通知quit
				fmt.Println(msg, "clean up")
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	quit := make(chan bool)
	c := boring("joe", quit)
	for i := 3; i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	close(quit)
	fmt.Println("main exit")
}
