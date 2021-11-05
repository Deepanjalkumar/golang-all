package main

import (
	"fmt"
)

func Input(Inputchan chan int, x int) {

	Inputchan <- x
}

func Intermediate(Inputchan chan int, Outputchan chan int) {

	s1 := <-Inputchan
	Outputchan <- s1
}

func Output(Outputchan chan int) {

	fmt.Println(<-Outputchan)
}

func main() {

	Inputchan := make(chan int)
	Outputchan := make(chan int)

	go Input(Inputchan, 2)
	go Intermediate(Inputchan, Outputchan)
	Output(Outputchan)

}
