package main

import (
	"fmt"
	"os"
	"runtime"
)

func drawCard(x int) string {
	switch x {
	case 1:
		return "Spades"
	case 2:
		return "Hearts"
	case 3:
		return "Diamonds"
	case 4:
		return "Clubs"
	default:
		return "Nonsence"
	}
}

func printStack() {
	var buf [4094]byte
	n := runtime.Stack(buf[:], false)
	f, _ := os.Create("stack.txt")
	f.Write(buf[:n])
	f.Close()
}

func main() {
	defer printStack()
	switch s := drawCard(2); s {
	case "Spades":
		fmt.Println("Spades")
	default:
		panic(fmt.Sprintf("invalid suit %q", s))
	}
}
