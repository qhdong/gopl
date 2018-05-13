package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	n := 0
	for input.Scan() {
		counts[input.Text()]++
		n++
	}

	fmt.Print("Word\tCount\n")
	fmt.Print("=====================\n")
	for word, cnt := range counts {
		fmt.Printf("%s\t%.2f%%\n", word, 100*float64(cnt)/float64(n))
	}
}
