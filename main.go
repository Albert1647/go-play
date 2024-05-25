package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	done := make(chan bool)
	// 	dones := make([]chan bool, 4)
	// dones[0] = make(chan bool)
	// 	dones[1] = make(chan bool)
	// 	dones[2] = make(chan bool)
	// 	dones[3] = make(chan bool)
	go greet("Nice to Meet You!", done)
	go greet("How are you!", done)
	go slowGreet("How.. Are.. You..", done)
	go greet("Hope you have a great day", done)
	for range done {
		// fmt.Println(doneChan)
	}
}
