package main

import (
	"fmt"
	"strings"
	"sync"
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

func exampleOfWaiting() {
	times := 10

	channel := make(chan int)
	for i := 1; i <= times; i++ {
		go AVeryLongRequest(i, channel)
	}

	// Wait for all goroutines to finish
	waitingTime := 250
	time.Sleep(time.Duration(waitingTime) * time.Millisecond) // Adjust the duration as needed
	fmt.Println("Waiting time: %v ms", waitingTime)
}

func exampleOfUsingChannel() {
	times := 10

	channel := make(chan int)
	for i := 1; i <= times; i++ {
		go AVeryLongRequest(i, channel)
	}
	// Receiving values from the channel
	for i := 1; i <= times; i++ {
		<-channel // Receive from the channel
	}

	fmt.Println("Done")
}

func exampleOfUsingSelect() {
	// We’ll use select to await both of these values simultaneously, printing each one as it arrives.
	c1 := make(chan int)
	c2 := make(chan int)

	go AVeryLongRequest(1000, c1)
	go AVeryLongRequest(2000, c2)

	for {
		select {
		case firstChannel := <-c1:
			fmt.Printf("first channel finished: %v\n", firstChannel)
		case secondChannel := <-c2:
			fmt.Printf("first channel finished: %v\n", secondChannel)
		// timeout
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}

}

func exampleOfHowLongSelectTake() {
	// it is just loop to next iteration
	// the execution time is just the condition
	c := make(chan int)

	for i := 1; i < 10; i++ {
		go func() {
			duration := time.Duration(i) * time.Second
			time.Sleep(duration)
			c <- i
		}()
	}

	for {
		select {
		case msg, ok := <-c:
			if !ok {
				fmt.Println("Channel closed!")
				return
			}
			fmt.Printf("Received: %v seconds\n", msg)
		}
	}
}

var mtxrw sync.RWMutex = sync.RWMutex{}
var locks map[string]time.Time = make(map[string]time.Time)
var total int
var lockMaxDuration = time.Duration(5) * time.Second

func acquiredLock(k string) (releaseFunc func(), acquireError error) {
	mtxrw.Lock()
	defer mtxrw.Unlock()
	// check if existed
	if acquiredAt, existed := locks[k]; existed {
		if time.Since(acquiredAt) < lockMaxDuration {
			return nil, fmt.Errorf("lock already acquired")
		}
		// else
		fmt.Println("Lock released!!!")
		fmt.Println(strings.Repeat("-", 50))
		delete(locks, k)
	}
	// otherwise add lock
	locks[k] = time.Now()
	// Create release function
	releaseFunc = func() {
		mtxrw.Lock()
		defer mtxrw.Unlock()
		if _, existed := locks[k]; existed {
			delete(locks, k)
		}
	}
	return releaseFunc, nil
}

func backgroundCleanUpLock() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		// clean up
		for k, acquiredAt := range locks {
			if time.Since(acquiredAt) > time.Duration(10)*time.Second {
				delete(locks, k)
			}
		}
		fmt.Printf("Clean up lock\n")
	}
}

func cronHandler(received time.Time) {
	keys := []string{"a"}
	processId := time.Now().String()
	//fmt.Printf("%v: Cron handler at %v", processId, received)
	for _, k := range keys {
		go processingKey(processId, k)
	}
}

func processingKey(processId string, k string) {
	releaseFunc, err := acquiredLock(k)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer releaseFunc()
	fmt.Printf("Start processing for key=%v processId=%v\n", k, processId)
	time.Sleep(time.Duration(8) * time.Second)
	total += 1
	fmt.Printf("processId=%v - total %v\n", processId, total)
	fmt.Printf("Finished processing for key=%v\n at %v", k, time.Now())
}

func main() {

	// cron init
	go backgroundCleanUpLock()
	for {
		cronChannel := make(chan time.Time, 1)
		time.Sleep(1 * time.Second)
		cronChannel <- time.Now()

		// cron handler
		select {
		case received, ok := <-cronChannel:
			defer close(cronChannel)
			if !ok {
				fmt.Println("cron ended!")
				return
			}
			// cron handler
			go cronHandler(received)
		}
	}

	// example.MutexLockExample()
	// exampleOfHowLongSelectTake()
}
