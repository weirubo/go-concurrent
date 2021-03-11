package main

import (
	"fmt"
	"time"
)

// channel

// sender

// receiver

func firstCh() <-chan int {
	c := make(chan int, 1)
	go func() {
		time.Sleep(time.Millisecond * 200)
		c <- 1
	}()
	return c
}

func secondCh() <-chan int {
	c := make(chan int, 1)
	go func() {
		time.Sleep(time.Millisecond * 300)
		c <- 2
	}()
	return c
}

func main() {
	select {
	case val, ok := <-firstCh():
		if !ok {
			fmt.Println("channel is closed")
			break
		}
		fmt.Println("first channel success", val)
	case val, ok := <-secondCh():
		if !ok {
			fmt.Println("channel is closed")
			break
		}
		fmt.Println("second channel success", val)
		// case <-time.After(time.Millisecond * 100):
		// 	fmt.Println("time out")
		// default:
		// 	fmt.Println("default")
	}
}
