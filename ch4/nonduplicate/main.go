package main

import "fmt"

func main() {
	s := []string{"a", "a", "b", "a", "a", "c", "c", "c"}
	fmt.Println(s)
	fmt.Println(nonDuplicate(s))
}

func nonDuplicate(strings []string) []string {
	if len(strings) <= 1 {
		return strings
	}
	out := strings[:1]
	pre := strings[0]
	for _, s := range strings[1:] {
		if s == pre {
			continue
		}
		out = append(out, s)
		pre = s
	}
	return out
}
