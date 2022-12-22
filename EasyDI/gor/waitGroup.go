package gor

import (
	. "EasyDI/reg"
	"sync"
	"time"
)

//var waitGroup = sync.WaitGroup{}

func doSum(count int, val *int, wGroup *sync.WaitGroup) {
	time.Sleep(time.Second)
	for i := 0; i < count; i++ {
		*val++
	}
	wGroup.Done()
}

func DoWait() {
	numRoutines := 3
	waitGroup := sync.WaitGroup{}
	counter := 0
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go doSum(5000, &counter, &waitGroup)
	}
	waitGroup.Wait()
	Printfln("Total: %v", counter)
}

func doSumMutex(count int, val *int, mutex *sync.Mutex, wGroup *sync.WaitGroup) {
	time.Sleep(time.Second)
	for i := 0; i < count; i++ {
		mutex.Lock()
		*val++
		mutex.Unlock()
	}
	wGroup.Done()
}

func DoMutex() {
	counter := 0
	numRoutines := 3
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(numRoutines)
	mutex := sync.Mutex{}
	for i := 0; i < numRoutines; i++ {
		go doSumMutex(5000, &counter, &mutex, &waitGroup)
	}
	waitGroup.Wait()
	Printfln("Mutex Total: %v", counter)

}
