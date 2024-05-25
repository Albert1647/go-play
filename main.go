package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {

	done := make(chan bool)
	go slowGreet("How.. Are.. You..", done)
	<-done
	// 	go greet("Nice to Meet You!")
	// 	go greet("How are you!")
	// 	go greet("Hope you have a great day")
}
