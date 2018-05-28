package main

import "fmt"

func join(sep string, a ...string) string {
	var ret string
	first := true
	for _, s := range a {
		if first {
			first = false
			ret += s
		} else {
			ret += sep + s
		}
	}
	return ret
}

func main() {
	names := []string{"aaron", "jack", "tony", "lucy"}
	sep := " , "
	fmt.Println(join(sep, names...))
	fmt.Println(join(sep, "good"))
}
