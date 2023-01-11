package main

import (
	"fmt"
	"sync"
)

var m sync.Map

func main() {
	m.Store(1, 1)
	m.Store(2, 1)
	fmt.Println(fibonacci(1))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(6))
	fmt.Println(fibonacci(7))
	fmt.Println(fibonacci(8))
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(11))
	fmt.Println(fibonacci(12))
}

func fibonacci(n int) int {
	if n <= 0 {
		panic("n is too low")
	}
	if v, ok := m.Load(n); ok {
		return v.(int)
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
