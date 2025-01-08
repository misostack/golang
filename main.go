package main

import (
	"fmt"
	"time"
)

var a = 0

func AVeryLongRequest(i int, channel chan int) {
	latencyValue := 200 + i*10
	latency := time.Duration(latencyValue) * time.Millisecond
	time.Sleep(latency)
	fmt.Printf("%v - Latency: %v ms . %v\n", i, latencyValue, time.Now())
	channel <- i
}

func main() {
	times := 10

	channel := make(chan int)
	for i := 1; i <= times; i++ {
		go AVeryLongRequest(i, channel)
	}

	// Wait for all goroutines to finish
	// waitingTime := 250
	// time.Sleep(time.Duration(waitingTime) * time.Millisecond) // Adjust the duration as needed
	// fmt.Println("Waiting time: %v ms", waitingTime)

	// Receiving values from the channel
	for i := 1; i <= times; i++ {
		<-channel // Receive from the channel
	}

	fmt.Println("Done")
}
