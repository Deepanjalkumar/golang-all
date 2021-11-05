package main

import (
	"fmt"
)

func Display(channel chan<- int) {

	channel <- 2

	close(channel)

}

func main() {

	channel := make(chan int)

	go Display(channel)

	fmt.Println(<-channel)

}
