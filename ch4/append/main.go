package main

import (
	"fmt"
)

func main() {
	var x []int
	for i := 0; i < 10; i++ {
		x = appendInt(x, i)
		fmt.Printf("%d  cap=%2d\t%v\n", i, cap(x), x)
	}
	x = appendInt(x, x...)
	fmt.Println(x)
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
