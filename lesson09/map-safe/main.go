package main

import (
	"fmt"
	"sync"
	"time"
)

// 构建并发安全的 map
type safeMap struct {
	m map[int]int
	sync.RWMutex
}

// 构造函数
func newSafeMap() *safeMap {
	sm := new(safeMap)
	sm.m = make(map[int]int)
	return sm
}

// 写
func (s *safeMap) store(key, val int) {
	s.Lock()
	s.m[key] = val
	s.Unlock()
}

// 读
func (s *safeMap) load(key int) int {
	s.RLock()
	val := s.m[key]
	s.RUnlock()
	return val
}

func main() {
	sm := newSafeMap()
	for i := 0; i < 10; i++ {
		go sm.store(i, i)
		go sm.load(i)
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println(sm.m)
}
