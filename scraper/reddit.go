package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// Scrape crawls and retrieves Gfycat / Imgur links from a list of subreddits
func Scrape() {
	c := colly.NewCollector(
		colly.AllowedDomains("old.reddit.com", "gfycat.com", "i.imgur.com"),
		colly.MaxDepth(2),
		colly.AllowURLRevisit(),
		colly.IgnoreRobotsTxt(),
		colly.Async(true),
	)
	extensions.RandomUserAgent(c)

	// Bypass any NSFW subreddits the crawler encounters
	c.OnHTML("form.pretty-form", func(e *colly.HTMLElement) {
		c.Post(e.Request.URL.String(), map[string]string{"over18": "yes"})
	})

	// Locate gifs/videos on Subreddit
	c.OnHTML(".top-matter", func(e *colly.HTMLElement) {
		url := e.ChildAttr("a[data-event-action=title]", "href")
		if isGfycat(url) {
			c.Visit(url)
		} else if isImgur(url) {
			// open url in browser
			fmt.Println(url)
		}
	})

	// Access GFYcat video source
	c.OnHTML(".video-player-container", func(e *colly.HTMLElement) {
		videoSrc := e.ChildAttr("source", "src")
		fmt.Println(videoSrc)
	})

	// Continue to the next page on Subreddit
	c.OnHTML("span.next-button", func(e *colly.HTMLElement) {
		nxt := e.ChildAttr("a", "href")
		c.Visit(nxt)
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })

	// Enter subreddits here
	reddits := []string{"https://old.reddit.com/r/funny"}

	for _, reddit := range reddits {
		c.Visit(reddit)
	}

	c.Wait()
}

func isGfycat(link string) bool {
	return strings.Contains(link, "gfycat")
}

func isImgur(link string) bool {
	return strings.Contains(link, "imgur")
}
