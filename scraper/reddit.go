package scraper

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alexbotello/GoGiffy/browse"
	"github.com/chromedp/chromedp"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// Scrape crawls and retrieves Gfycat / Imgur links from a list of subreddits
func Scrape(ctxt context.Context, chrome *chromedp.CDP) {
	c := colly.NewCollector(
		colly.AllowedDomains("old.reddit.com", "gfycat.com", "i.imgur.com"),
		colly.MaxDepth(2),
		colly.AllowURLRevisit(),
		colly.IgnoreRobotsTxt(),
		// colly.Async(true),
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
			browse.GoTo(ctxt, chrome, url)
			time.Sleep(10 * time.Second)
		}
	})

	// Access GFYcat video source
	c.OnHTML(".video-player-container", func(e *colly.HTMLElement) {
		videoSrc := e.ChildAttr("source", "src")
		fmt.Println(videoSrc)
		browse.GoTo(ctxt, chrome, videoSrc)
		time.Sleep(10 * time.Second)
	})

	// Continue to the next page on Subreddit
	c.OnHTML("span.next-button", func(e *colly.HTMLElement) {
		nxt := e.ChildAttr("a", "href")
		c.Visit(nxt)
	})

	// Enter subreddit as CLI argument eg. `./GoGiffy https:/old.reddit.com/r/gifs`
	reddits := os.Args[1:]

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
