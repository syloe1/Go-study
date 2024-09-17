package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// Define a struct to represent a web page
type Page struct {
	URL  string
	Body string
}

// Define an interface for the crawler
type Crawler interface {
	Crawl(url string) (*Page, error)
}

// Concrete implementation of Crawler
type SimpleCrawler struct{}

func (sc *SimpleCrawler) Crawl(url string) (*Page, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &Page{
		URL:  url,
		Body: string(body),
	}, nil
}
func main() {
	urls := []string{
		"http://example.com",
		"http://example.org",
		"http://example.net",
	}

	var wg sync.WaitGroup
	results := make(chan *Page)

	crawler := &SimpleCrawler{}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			page, err := crawler.Crawl(url)
			if err != nil {
				log.Printf("Error crawling %s: %v", url, err)
				return
			}
			results <- page
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for page := range results {
		fmt.Printf("URL: %s\n", page.URL)
		fmt.Printf("Content: %s\n", page.Body[:100]) // Print first 100 characters of the body
	}
}
