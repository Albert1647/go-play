package main

import "fmt"

func main() {
	output("Hello World!")
}

func output(array ...interface{}) {
	fmt.Println(array...)
}
