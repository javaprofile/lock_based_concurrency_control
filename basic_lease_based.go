package main

import (
	"fmt"
	"sync"
	"time"
)

type LeaseLock struct {
	mu        sync.Mutex
	owner     string
	expiry    time.Time
	leaseTime time.Duration
}

func (l *LeaseLock) AcquireLock(clientID string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if time.Now().After(l.expiry) || l.owner == "" {
		l.owner = clientID
		l.expiry = time.Now().Add(l.leaseTime)
		return true
	}
	return false
}

func (l *LeaseLock) ReleaseLock(clientID string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.owner == clientID {
		l.owner = ""
		return true
	}
	return false
}

func main() {
	lock := &LeaseLock{leaseTime: 2 * time.Second}
	
	client1 := "Client1"
	client2 := "Client2"
	
	if lock.AcquireLock(client1) {
		fmt.Println(client1, "acquired the lock")
		time.Sleep(1 * time.Second)
		if lock.ReleaseLock(client1) {
			fmt.Println(client1, "released the lock")
		}
	} else {
		fmt.Println(client1, "could not acquire the lock")
	}
	
	if lock.AcquireLock(client2) {
		fmt.Println(client2, "acquired the lock")
		time.Sleep(3 * time.Second)
		if lock.ReleaseLock(client2) {
			fmt.Println(client2, "released the lock")
		}
	} else {
		fmt.Println(client2, "could not acquire the lock")
	}
}
