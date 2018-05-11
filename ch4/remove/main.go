package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// fmt.Println(remove(s, 3))
	fmt.Println(remove2(s, 3))
}

func remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func remove2(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
