package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/Carmo-sousa/html"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - Canal somente leitura

func titulo(urls ...string) <-chan string {
	// Pattern: `[^<title>]\w+[^<\/title>]`gm
	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile(`<title>(.*?)<\/title>`)
			ch <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return ch
}

func main() {
	t1 := html.Titulo("https://www.google.com.br", "https://github.com/")
	t2 := titulo("https://www.google.com.br", "https://github.com/")
	fmt.Println(<-t1)
	fmt.Println(<-t1)
	fmt.Println(<-t2)
	fmt.Println(<-t2)
}
