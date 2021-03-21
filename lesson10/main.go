package main

import (
	"fmt"
	"runtime"
	"sync"
)

// sync.Pool

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("New 一个新对象")
			return 0
		},
	}

	// 取，临时对象池中没有数据，会调用 New，New 创建一个新对象直接返回，不会存储在临时对象池中
	val := pool.Get().(int)
	fmt.Println(val)

	// 存
	pool.Put(10)
	// 手动调用 GC()，用于验证 GC 之后，临时对象池中的对象会被清空。
	runtime.GC()

	// 取
	val2 := pool.Get().(int)
	fmt.Println(val2)
}
