package main

import (
	"fmt"
	"sync"
	"time"
)

type MutexCounter struct {
	mu sync.Mutex
	mp map[string]int
}

func (m *MutexCounter) Incr(in string) {
	m.mu.Lock()
	m.mp[in]++
	m.mu.Unlock()
}

func (m *MutexCounter) Value(in string) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.mp[in]
}

func main() {
	mc := MutexCounter{mp: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go mc.Incr("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(mc.Value("somekey"))
}
