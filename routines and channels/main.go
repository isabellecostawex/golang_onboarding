package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("worker:", workerId, "Msg:", res)
		time.Sleep(time.Second)
	}
}

func main() {
	msg := make(chan int)

	go worker(1, msg)
	go worker(2, msg)
	go worker(3, msg)
	go worker(4, msg)
	go worker(5, msg)

	for i := 0; i < 10; i++ {
		msg <- i
	}
	time.Sleep(5*time.Second)
}
