package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	// Base URL of the website to restrict crawling
	baseURL := "http://testphp.vulnweb.com"

	// A set to keep track of visited URLs
	visited := make(map[string]bool)

	// Recursive function to crawl the website
	var crawl func(url string)
	crawl = func(url string) {
		if visited[url] {
			return
		}
		visited[url] = true

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")

			// Normalize the link by resolving it relative to the base URL
			absoluteURL := e.Request.AbsoluteURL(link)

			// Check if the link is a full URL (with scheme and domain) and starts with the base URL
			if strings.HasPrefix(absoluteURL, baseURL) {
				fmt.Println(absoluteURL)
				crawl(absoluteURL) // Recursively crawl the found URL
			}
		})

		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "Error:", err)
		})

		// Visit the URL
		err := c.Visit(url)
		if err != nil {
			fmt.Println("Error while visiting URL:", url, "Error:", err)
		}
	}

	// Read URLs from standard input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		crawl(url)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}
