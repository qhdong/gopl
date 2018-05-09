package main

import "fmt"

func main() {
	fmt.Println(gcd(30, 293))
	fmt.Println(fibonacci(30))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fibonacci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
