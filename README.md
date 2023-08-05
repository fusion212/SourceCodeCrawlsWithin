# SourceCodeCrawlsWithin
  
  # Crawl

  Crawl is a simple web crawler written in Go that recursively crawls a website and prints all the URLs it finds. It uses the `gocolly` library for web scraping and URL parsing.

  ## Features

  - Recursively crawls a website and prints all the URLs it finds.
  - Restricts crawling to a specific base URL to prevent crawling outside the target website.
  - Handles relative and absolute URLs.

  ## Installation

  1. Make sure you have Go installed on your system. If not, you can download it from the official website: [https://golang.org/dl/](https://golang.org/dl/)

  2. Clone this repository to your local machine:

     ```
     git clone https://github.com/yourusername/crawl.git
     ```

  3. Change into the project directory:

     ```
     cd crawl
     ```

  4. Build the executable binary:

     ```
     go build crawl.go
     ```

  ## Usage

  You can use the "crawl" tool in two ways:

  1. Provide URLs as command-line arguments:

     ```
     ./crawl http://example.com https://example.com/page1
     ```

  2. Pipe URLs from standard input:

     ```
     echo -e 'http://example.com\nhttps://example.com/page1' | ./crawl
     ```

  The tool will recursively crawl each URL and print all the URLs it finds on the same domain.

