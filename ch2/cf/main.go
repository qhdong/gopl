package main

import (
	"fmt"
	"github.com/qhdong/gopl/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		c, f := tempconv.Celsius(t), tempconv.Fahrenheit(t)
		fmt.Printf("%s = %s, %s = %s\n",
			c, tempconv.CToF(c), f, tempconv.FToC(f))
	}
}
