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
