package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 7}
	// 结果应该是 5, 1, 2, 4, 7
	s = Add(s, 0, 5)
	fmt.Println(s, 5, 1, 2, 4, 7)

	// 结果应该是5, 9, 1, 2, 4, 7
	s = Add(s, 1, 9)
	fmt.Println(s, 5, 9, 1, 2, 4, 7)

	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s = Add(s, 6, 13)
	fmt.Println(s, 9, 1, 2, 4, 7, 13)

	// 结果应该是5, 9, 2, 4, 7, 13
	s = Delete(s, 2)
	fmt.Println(s, 5, 9, 2, 4, 7, 13)

	// 结果应该是9, 2, 4, 7, 13
	s = Delete(s, 0)
	fmt.Println(s, 9, 2, 4, 7, 13)

	// 结果应该是9, 2, 4, 7
	s = Delete(s, 4)
	fmt.Println(s, 9, 2, 4, 7)

}

func Add(s []int, index int, value int) []int {
	if index < 0 {
		panic("index lower then zero")
	}
	s = append(s, 0)
	for i := len(s) - 1; i >= 0; i-- {
		if index >= i {
			s[i] = value
			break
		}
		s[i] = s[i-1]
	}
	return s
}

func Delete(s []int, index int) []int {
	if index < 0 {
		panic("index lower then zero")
	}
	for i := 0; i < len(s); i++ {
		if index < i {
			s[i-1] = s[i]
		}
		if i == len(s)-1 && index <= i {
			s = s[:len(s)-1]
		}
	}
	return s
}
