package main

import (
	"flag"
	"fmt"
	"github.com/wsxiaoys/terminal/color"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type CheckResult struct {
	idx      int
	url      string
	status   bool
	err      bool
	duration *time.Duration
}

func main() {
	var quiet bool = false
	var result bool = true
	var results []*CheckResult

	flag.Usage = func() {
		fmt.Printf("Usage: testgzip [options] URL+\n\nOptions:\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&quiet, "quiet", false, "Disables any output and limits the status to the exit code")
	flag.Parse()

	if 0 == len(flag.Args()) {
		fmt.Println("You have to provide at least one URL to be tested.")
		os.Exit(2)
	}

	// First, let's make sure that all passed arguments are actually valid URLs.
	for _, url_ := range flag.Args() {
		if !isUrl(url_) {
			fmt.Printf("%s is not a valid URL\n", url_)
			os.Exit(1)
			return
		}
	}

	results = make([]*CheckResult, len(flag.Args()))

	resultChan := make(chan CheckResult, len(flag.Args()))

	for i, url_ := range flag.Args() {
		go testUrl(url_, i, resultChan)
	}

	for i := 0; i < len(flag.Args()); i++ {
		checkResult := <-resultChan
		results[checkResult.idx] = &checkResult
		if checkResult.err || !checkResult.status {
			result = false
		}
	}

	if !quiet {
		for _, checkResult := range results {
			if checkResult.err {
				color.Printf("@{r}ERROR@{|}  %s\n", checkResult.url)
			} else {
				if checkResult.status {
					color.Printf("@{g}OK@{|}     %s (%s)\n", checkResult.url, checkResult.duration)
				} else {
					color.Printf("@{r}FAILED@{|} %s (%s)\n", checkResult.url, checkResult.duration)
				}
			}
		}
	}

	if !result {
		os.Exit(-1)
	} else {
		os.Exit(0)
	}
}

// Checks if a given string represents a URL that we can work with.
func isUrl(u string) bool {
	result, err := url.Parse(u)
	if err != nil {
		return false
	}
	return result.Scheme == "http" || result.Scheme == "https"
}

// This method tests if the given URL responds in a gzipped manner.
func testUrl(url string, idx int, result chan CheckResult) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Failed to generate new request")
		result <- CheckResult{idx, url, false, true, nil}
		return
	}
	req.Header.Add("Accept-Encoding", "gzip,deflate")
	start := time.Now()
	resp, err := client.Do(req)
	duration := time.Since(start)
	if err != nil {
		log.Println("Request failed: %s", err)
		result <- CheckResult{idx, url, false, true, &duration}
		return
	}
	encoding := resp.Header.Get("Content-Encoding")
	if strings.Contains(encoding, "gzip") || strings.Contains(encoding, "deflate") {
		result <- CheckResult{idx, url, true, false, &duration}
		return
	}
	result <- CheckResult{idx, url, false, false, &duration}
}
