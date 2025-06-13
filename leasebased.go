package main
import (
	"fmt"
	"sync"
	"time"
)
type LeaseLock struct {
	mu            sync.Mutex
	isLocked      bool
	lockHolder    string
	lockExpiration time.Time
	lockDuration  time.Duration
}
func NewLeaseLock(duration time.Duration) *LeaseLock {
	return &LeaseLock{
		lockDuration: duration,
	}
}
func (l *LeaseLock) AcquireLock(holder string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.isLocked || time.Now().After(l.lockExpiration) {
		l.isLocked = true
		l.lockHolder = holder
		l.lockExpiration = time.Now().Add(l.lockDuration)
		return true
	}
	return false
}
func (l *LeaseLock) ReleaseLock() {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.isLocked {
		l.isLocked = false
		l.lockHolder = ""
	}
}
func (l *LeaseLock) IsLocked() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.isLocked
}
func (l *LeaseLock) GetLockHolder() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.lockHolder
}
func simulateLocking(lock *LeaseLock, client string, lockCh chan bool) {
	success := lock.AcquireLock(client)
	if success {
		lockCh <- true
		time.Sleep(lock.lockDuration)
		lock.ReleaseLock()
	} else {
		lockCh <- false
	}
}
func trackLocksPerSecond(lock *LeaseLock, duration time.Duration) {
	var lockCount int
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("Locks acquired in last second: %d\n", lockCount)
			lockCount = 0
		}
	}
}
func main() {
	lock := NewLeaseLock(2 * time.Second)
	lockCh := make(chan bool)

	go trackLocksPerSecond(lock, 1*time.Second)
	go simulateLocking(lock, "Client1", lockCh)
	go simulateLocking(lock, "Client2", lockCh)
	go simulateLocking(lock, "Client3", lockCh)
	for i := 0; i < 3; i++ {
		if <-lockCh {
			fmt.Println("Lock acquired successfully!")
		} else {
			fmt.Println("Failed to acquire lock.")
		}
	}
	time.Sleep(5 * time.Second)
}
