package main

import (
	"fmt"
	"time"
)

func main() {
	req := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		req <- i
	}

	close(req)

	limiter := time.Tick(time.Millisecond * 200)

	for r := range req {
		<-limiter
		fmt.Println("Request", r, time.Now())
	}

	burstLimit := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstLimit <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstLimit <- t
		}
	}()

	burstReq := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstReq <- i
	}
	close(burstReq)

	for br := range burstReq {
		<-burstLimit
		fmt.Println("Req: ", br, time.Now())
	}
}
