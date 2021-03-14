package main

import (
	"fmt"
	"time"
)

// golang 原生 map

var m = make(map[int]int)

func main() {
	for i := 0; i < 3; i++ {
		go store(i, i)
		go load(i)
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println(m)
}

func store(key, val int) {
	m[key] = val
}

func load(key int) int {
	return m[key]
}
