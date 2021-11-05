package main

import (
	"fmt"
)

func main() {

	channel := make(chan int, 2)

	channel <- 1
	channel <- 2

	fmt.Println(<-channel)
	fmt.Println(<-channel)

}
