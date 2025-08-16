# Golang Cheatsheet

## Concurrency

### Explanation

- A goroutine is a lightweight thread managed by the Go runtime.

```sh
# register routine
go f(x, y, z)
# start routine
f(x, y, z)
```

These code will run in sequence

```go
var a = 0

func AVeryLongRequest(i int) {
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("%v . %v\n", i, time.Now())
}

func main() {
	times := 10
	for i := 1; i <= times; i++ {
		AVeryLongRequest(i)
	}
}
```

So that we wanna speed up by run them in parallel

````go
package main

import (
	"fmt"
	"time"
)

var a = 0

func AVeryLongRequest(i int) {
	latencyValue := 200 + i*10
	latency := time.Duration(latencyValue) * time.Millisecond
	time.Sleep(latency)
	fmt.Printf("%v - Latency: %v ms . %v\n", i, latencyValue, time.Now())
}

func main() {
	times := 10
	waitingTime := 250
	for i := 1; i <= times; i++ {
		go AVeryLongRequest(i)
	}

	// Wait for all goroutines to finish
	time.Sleep(time.Duration(waitingTime) * time.Millisecond) // Adjust the duration as needed
	fmt.Println("Waiting time: %v ms", waitingTime)
}

```sh
1 - Latency: 210 ms . 2025-01-08 16:53:28.464612 +0700 +07 m=+0.211178418
2 - Latency: 220 ms . 2025-01-08 16:53:28.474615 +0700 +07 m=+0.221180709
3 - Latency: 230 ms . 2025-01-08 16:53:28.484623 +0700 +07 m=+0.231189126
4 - Latency: 240 ms . 2025-01-08 16:53:28.494284 +0700 +07 m=+0.240849793
5 - Latency: 250 ms . 2025-01-08 16:53:28.504582 +0700 +07 m=+0.251148293
Waiting time: %v ms 250
````

=> how about other request, they disappear

### Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

```go
// (The data flows in the direction of the arrow.)
// Like maps and slices, channels must be created before use:
ch := make(chan int)
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.

```

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

```go
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

```

**Note**: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic

### With Mutex and RWMutex

```go
package example

import (
	"fmt"
	"sync"
	"time"
)

// let's say we have a params can be in memory or redis
// in this case we use memory
var a int
var mtx = sync.Mutex{}
var mtxrw = sync.RWMutex{}

func increaseValue() {
	mtx.Lock()
	defer mtx.Unlock()
	a++
}

func readValue(i int) {
	fmt.Printf("%v. Current a value = %v\n", i, a)
}

func increaseValueLockWithRWMutex() {
	// allow read
	mtxrw.Lock()
	defer mtxrw.Unlock()
	a++
}

func MutexLockExample() {

	times := 500
	for i := 0; i < times; i++ {
		//go increaseValue()
		go increaseValueLockWithRWMutex()
	}
	for i := 1; i <= 500; i++ {
		go readValue(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(a)
}
```

### Decimal

```sh
go get github.com/shopspring/decimal
```

### References

- https://go.dev/wiki/MutexOrChannel
- https://go.dev/tour/concurrency/1
- https://medium.com/@asgrr/golang-sync-4787b18fee41
- https://viblo.asia/p/go-sync-package-6-khai-niem-chinh-ban-can-biet-Ny0VGj2pLPA
- https://psj.codes/concurrency-in-go-part-1-goroutines-channels-and-select

## Target for 1st Quarter of 2025

Go through all the basic topics at https://go.dev/doc/tutorial/
