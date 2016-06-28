package main

import "fmt"

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

}
