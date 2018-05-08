package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		log.Printf("Get %s, Status: %s\n", url, resp.Status)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
