package main

import "fmt"

func ArrayContains[K comparable](arr []K, key K) bool {
	for _, k := range arr {
		if k == key {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(ArrayContains([]int{1, 2, 3}, 5))
	fmt.Println(ArrayContains([]int{1, 2, 3}, 1))
	fmt.Println(ArrayContains([]string{"a", "b", "c"}, "a"))
	fmt.Println(ArrayContains([]string{"a", "b", "c"}, "d"))
}
