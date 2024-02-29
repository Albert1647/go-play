package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

func output(array ...interface{}) {
	fmt.Println(array...)
}
