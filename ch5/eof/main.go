package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		_, _, err := in.ReadRune()
		if err == io.EOF {
			log.Println("Reach end")
			break
		}
		if err != nil {
			log.Fatalf("read failed: %v", err)
		}
	}
}
