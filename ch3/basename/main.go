package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "a/b/c.d.go"
	fmt.Printf("basename1: %s -> %s\n", path, basename1(path))
	fmt.Printf("basename2: %s -> %s\n", path, basename2(path))
}

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
