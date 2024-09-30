package main

import (
	"fmt"
	"time"
)

func main() {

	// Ref := https://www.freecodecamp.org/news/concurrent-programming-in-go
	msg := make(chan string)
	go greet(msg)
	time.Sleep(1 * time.Second)

	greeting := <-msg

	time.Sleep(2 * time.Second)
	fmt.Println("Received the greeting: ", greeting)

}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!")

	ch <- "Hello Rwitesh"

	fmt.Println("Greeter completed!")
}
