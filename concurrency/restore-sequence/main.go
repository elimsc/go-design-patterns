package main

import (
	"fmt"
	"math/rand"
	"time"
)

// merge N channels into 1 channel
func fanIn(cs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for _, ci := range cs {
		go func(cc <-chan Message) {
			for v := range cc { // 当cc被close时, 这个goroutine退出
				c <- v
			}
		}(ci)
	}
	return c
}

type Message struct {
	msg  string
	wait chan bool
}

// 返回一个带内部channel的msg channel, 调用者可以通过这个内部的channel控制goroutine的流程
func boring(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			msg := Message{msg: fmt.Sprintf("%s %d", msg, i), wait: waitForIt}
			c <- msg
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			<-waitForIt
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ane"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.msg)
		msg2 := <-c
		fmt.Println(msg2.msg)

		// 让boring goroutine继续
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("main exit.")
}
