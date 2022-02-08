package main

import (
	"fmt"
)

func main() {
    // create a channel 
	ch := make(chan int)

//create a goroutine sender
	go sender(ch)
   // send a value to channel
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40

	close(ch)

}
func sender(ch chan int) {
	
//iterate all the channel value
	for i := range ch {
		fmt.Println(i)

	}

}
