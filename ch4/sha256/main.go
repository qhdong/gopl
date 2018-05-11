package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/qhdong/gopl/ch2/popcount"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%#x\n%#x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	cnt := 0
	for i := 0; i < len(c1); i++ {
		cnt += popcount.PopCount(uint64(c1[i] ^ c2[i]))
	}
	fmt.Printf("There are %d(%.2f%%) bits different\n", cnt, 100*float64(cnt)/256)

	zero2(&c1)
	fmt.Printf("%#x\n", c1)
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero2(ptr *[32]byte) {
	*ptr = [32]byte{}
}
