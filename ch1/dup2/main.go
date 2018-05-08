// dup2 prints duplicate lines, it process both stdin and argument files
package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    counts := make(map[string]int)
    if len(os.Args[1:]) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range os.Args[1:] {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
}
