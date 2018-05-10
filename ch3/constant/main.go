package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
)

func main() {
	fmt.Printf("MiB = %d\n", MiB)
}
