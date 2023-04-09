package main

import (
	"fmt"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		// scrapes pages through concurrency
		go func(pageNum int) {
			defer wg.Done()
			scrapeHackerNewsPosts(pageNum)
		}(i)
	}

	wg.Wait()
}

func scrapeHackerNewsPosts(pageNum int) {
	url := fmt.Sprintf("https://news.ycombinator.com/news?p=%d", pageNum)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	doc.Find(".title a").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		fmt.Printf("Post %d: %s\n", index+1, title)
	})
}
