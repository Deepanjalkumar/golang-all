package main

import (
	"fmt"
)

func Display(channel chan<- int) {

	for i := 1; i <= 5; i++ {
		channel <- i
	}

	close(channel)
}

func main() {

	channel := make(chan int)

	go Display(channel)

	for value := range channel {
		fmt.Println(value)
	}
}
