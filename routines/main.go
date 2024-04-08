package main

import (
	"fmt"
	"time"
)

func counter(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main() {
	go counter("a")
	go counter("b")
	time.Sleep(time.Second * 10)

}
