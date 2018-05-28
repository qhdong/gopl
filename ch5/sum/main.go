package main

import "fmt"

func min(x int, vals ...int) int {
	for _, val := range vals {
		if x > val {
			x = val
		}
	}
	return x
}

func max(x int, vals ...int) int {
	for _, val := range vals {
		if x < val {
			x = val
		}
	}
	return x
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	vals := []int{1, 2, 3}
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(vals...))

	fmt.Println(max(1))
	fmt.Println(max(1, 2, 3))
}
