package main

import "fmt"
import "strconv"

func main() {
	var x uint8 = 1<<1 | 1<<3
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	xx := int64(0xdeadbeef)
	fmt.Printf("%d %#[1]x %#[1]X\n", xx)

	unicode := '😔'
	fmt.Printf("%d %[1]c %[1]q\n", unicode)

	fmt.Println(strconv.Itoa(23))
	fmt.Println(strconv.Atoi("23"))
	fmt.Println(strconv.ParseInt("10101", 2, 0))
}
