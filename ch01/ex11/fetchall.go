package main

import (
	"fmt"
	"io"
//	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
)

const HTTP_PROTOCOL_INDICATER = "http://"
const HTTPS_PROTOCOL_INDICATER = "https://"


func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	domain :=  "dammy_filename"

	if strings.HasPrefix(url, HTTP_PROTOCOL_INDICATER) {
		domain = strings.TrimLeft(url, HTTP_PROTOCOL_INDICATER)
	} else if strings.HasPrefix(url, HTTPS_PROTOCOL_INDICATER) {
		domain = strings.TrimLeft(url, HTTPS_PROTOCOL_INDICATER)
	} else {
		domain = url
	}

	dstFile, err := os.Create(domain)
	if err != nil {
		ch <- fmt.Sprintf("open file error:  %s", domain)
		return
	}

	defer dstFile.Close() // File close when function fetch is closed

	//nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	nbytes, err := io.Copy(dstFile, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
