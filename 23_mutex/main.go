

//mutex is used to protect shared data from being accessed by multiple goroutines at the same time
//this is important to prevent data	aces and ensure data consistency

//a mutex has two states: locked and unlocked
//when a goroutine locks a mutex, other goroutines that try to lock the same mutex will block until the mutex is unlocked
//this ensures that only one goroutine can access the shared data at a time

//to use a mutex, we need to import the sync package
//we can create a mutex using sync.Mutex{}
//we can lock a mutex using the Lock() method and unlock it using the Unlock() method

//it is important to always unlock a mutex after locking it to prevent deadlocks
//a deadlock occurs when two or more goroutines are waiting for each other to release a resource, causing them to be stuck indefinitely

package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)

	}

	wg.Wait()
	fmt.Println(myPost.views)
}