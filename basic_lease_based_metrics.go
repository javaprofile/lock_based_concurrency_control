package main

import (
	"fmt"
	"sync"
	"time"
)

type LockTracker struct {
	mu       sync.Mutex
	count    int
	lastTime time.Time
}

func (lt *LockTracker) Increment() {
	lt.mu.Lock()
	lt.count++
	lt.mu.Unlock()
}

func (lt *LockTracker) Reset() int {
	lt.mu.Lock()
	defer lt.mu.Unlock()
	locks := lt.count
	lt.count = 0
	lt.lastTime = time.Now()
	return locks
}

func main() {
	tracker := &LockTracker{}
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			tracker.Increment()
		}
	}()
	for {
		select {
		case <-ticker.C:
			locksPerSec := tracker.Reset()
			fmt.Printf("Locks per second: %d\n", locksPerSec)
		}
	}
}
