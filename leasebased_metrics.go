package main

import (
	"fmt"
	"sync"
	"time"
)

type LockTracker struct {
	mu        sync.Mutex
	lockCount int
}

func (lt *LockTracker) Increment() {
	lt.mu.Lock()
	lt.lockCount++
	lt.mu.Unlock()
}

func (lt *LockTracker) Reset() int {
	lt.mu.Lock()
	count := lt.lockCount
	lt.lockCount = 0
	lt.mu.Unlock()
	return count
}

func main() {
	tracker := &LockTracker{}
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			tracker.Increment()
		}
	}()

	for range ticker.C {
		fmt.Printf("Locks/sec: %d\n", tracker.Reset())
	}
}
