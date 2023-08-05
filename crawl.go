package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	// A set to keep track of visited URLs
	visited := make(map[string]bool)

	// Recursive function to crawl the website
	var crawl func(u string)
	crawl = func(u string) {
		if visited[u] {
			return
		}
		visited[u] = true

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")

			// Normalize the link by resolving it relative to the base URL
			absoluteURL := e.Request.AbsoluteURL(link)

			// Check if the link is a full URL (with scheme and domain)
			parsedURL, err := url.Parse(absoluteURL)
			if err == nil && parsedURL.Scheme != "" && parsedURL.Host != "" {
				fmt.Println(absoluteURL)
				crawl(absoluteURL) // Recursively crawl the found URL
			}
		})

		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "Error:", err)
		})

		// Visit the URL
		err := c.Visit(u)
		if err != nil {
			fmt.Println("Error while visiting URL:", u, "Error:", err)
		}
	}

	// Read URLs from standard input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		urlInput := scanner.Text()
		crawl(urlInput)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}
