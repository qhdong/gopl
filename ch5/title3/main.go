package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)

	if title == "" {
		err = fmt.Errorf("no title element")
	}
	return title, err
}

func getDoc(url string) (doc *html.Node, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	cf := resp.Header.Get("Content-Type")
	if cf != "text/html" && !strings.HasPrefix(cf, "text/html;") {
		err = fmt.Errorf("%s has type %s, not text/html", url, cf)
		return nil, err
	}

	doc, err = html.Parse(resp.Body)
	return
}

func main() {
	for _, url := range os.Args[1:] {
		doc, err := getDoc(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "soleTitle: %v\n", err)
			continue
		}
		title, err := soleTitle(doc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "soleTitle: %v", err)
		}
		fmt.Println(title)
	}
}
