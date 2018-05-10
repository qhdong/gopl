package main

import (
	"bytes"
	"fmt"
)

func main() {
	for _, s := range []string{"12", "123", "1234", "123934929348"} {
		fmt.Printf("comma: %s -> %s\n", s, comma(s))
		fmt.Printf("comma2: %s -> %s\n", s, comma2(s))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// 12345 -> 12,345
func comma2(s string) string {
	if len(s) <= 3 {
		return s
	}
	var buf bytes.Buffer
	idx := len(s) % 3
	buf.WriteString(s[:idx])
	for idx < len(s) {
		if idx > 0 {
			buf.WriteRune(',')
		}
		buf.WriteString(s[idx : idx+3])
		idx += 3
	}
	return buf.String()
}
