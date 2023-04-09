# WebScrapper-Go

A Web scrapper that utilizes go concurrency to scrap hacker news

Created through chatgpt help

Let's walk through this code step by step:

We import the necessary packages, including "fmt", "sync", and "github.com/PuerkitoBio/goquery".
We create a WaitGroup to wait for all goroutines to complete.
We use a for loop to spawn 10 goroutines, each of which will scrape a different page of Hacker News.
In each goroutine, we use the "goquery.NewDocument" function to retrieve the HTML document for the page, and the "doc.Find" function to select all the titles of the posts.
Finally, we use the "wg.Done" function to signal that the goroutine has completed, and the WaitGroup's "Wait" function to block until all goroutines have completed.
This example demonstrates how to use Go's concurrency features to speed up a web scraping task by processing multiple pages in parallel. By using goroutines to process each page independently, we can take advantage of the CPU's multiple cores and reduce the overall runtime of the program.

## To use the script

- `Go main.go`
