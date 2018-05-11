package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/sha3"
	"log"
	"os"
)

var n = flag.Int("crypt", 2, "加密级别：2->256, 3->384, 5->512")

func main() {
	flag.Parse()
	if *n != 2 && *n != 3 && *n != 5 {
		log.Fatal("-crypt参数只能是:2, 3, 5")
	}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		switch *n {
		case 2:
			fmt.Printf("%#X\n", sha3.Sum256([]byte(input.Text())))
		case 3:
			fmt.Printf("%#X\n", sha3.Sum384([]byte(input.Text())))
		case 5:
			fmt.Printf("%#X\n", sha3.Sum512([]byte(input.Text())))
		default:
			continue
		}
	}
}
