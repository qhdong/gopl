// echo3 prints command line arguments more effectively
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(strings.Join(os.Args[1:], ", "))
    fmt.Println(os.Args)
}
