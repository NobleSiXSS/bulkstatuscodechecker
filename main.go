package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func get(client http.Client, url string) int {
	resp, err := client.Get(url)
	if err != nil {
		log.Println("Error", url)
	}
	if resp == nil {
		return 0
	}
	return resp.StatusCode
}

func main() {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}

	var client = http.Client {
		Transport: transCfg,
		Timeout: time.Second * 7,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	var numJobs int
	flag.IntVar(&numJobs,"t", 20, "Number of threads")
	flag.Parse()

	urls := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < numJobs; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for url := range urls {
				status := get(client, url)
				if status != 0 {
					fmt.Println(status, url)
				}
			}
		}()
	}

	buffstdin := bufio.NewScanner(os.Stdin)
	for buffstdin.Scan() {
		url := strings.ToLower(buffstdin.Text())
		urls <- url
	}
	close(urls)
	wg.Wait()
}
