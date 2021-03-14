package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Map

func main() {
	for i := 0; i < 10; i++ {
		go store(i, i)
		go load(i)
	}
	time.Sleep(time.Millisecond * 100)
	m.Range(list)
}

func store(key, val int) {
	m.Store(key, val)
}

func load(key int) interface{} {
	if val, ok := m.Load(key); ok {
		return val
	}
	return nil
}

func list(key, val interface{}) bool {
	fmt.Printf("key:%d=>val:%d\n", key, val)
	return true
}
