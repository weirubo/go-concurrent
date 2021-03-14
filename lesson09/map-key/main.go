package main

import "fmt"

// key

func main() {
	m1 := map[interface{}]string{
		1:       "A",
		"2":     "B",
		[]int{}: "C",
	}
	fmt.Println(m1)
}
