package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "你好吗？👏"
	fmt.Printf("len(s) = %d, RuneCountInString(s) = %d\n", len(s), utf8.RuneCountInString(s))

	r := []rune(s)
	fmt.Println(r)

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
