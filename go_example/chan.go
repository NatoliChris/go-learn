package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)

	//checking out buffered channels

	bufferedmsgs := make(chan string, 2)

	bufferedmsgs <- "buffered"
	bufferedmsgs <- "channel"

	fmt.Println(<-bufferedmsgs)
	fmt.Println(<-bufferedmsgs)

	//learning about synchronisation of threads
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
