package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	s := a[:]
	fmt.Println(s)
	rev(a[:])
	fmt.Println(s)
	rotate(a[:], 3)
	fmt.Println(a)

	fmt.Println([]int(nil) == nil, make([]int, 3)[3:] == nil, []int{} == nil)
}

func rev(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, n int) {
	rev(s[n:])
	rev(s[:n])
	rev(s)
}
