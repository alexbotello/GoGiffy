package scraper

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alexbotello/GoGiffy/browse"
	"github.com/chromedp/chromedp"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// Scrape crawls and retrieves Gfycat / Imgur links from a list of subreddits
func Scrape(ctxt context.Context, chrome *chromedp.CDP) {
	duration := getDuration()

	c := colly.NewCollector(
		colly.AllowedDomains("old.reddit.com", "gfycat.com", "i.imgur.com", "redd.it"),
		colly.MaxDepth(2),
		colly.AllowURLRevisit(),
		colly.IgnoreRobotsTxt(),
	)
	extensions.RandomUserAgent(c)

	// Bypass any NSFW subreddits the crawler encounters
	c.OnHTML("form.pretty-form", func(e *colly.HTMLElement) {
		c.Post(e.Request.URL.String(), map[string]string{"over18": "yes"})
	})

	// Locate gifs/videos on Subreddit
	c.OnHTML(".top-matter", func(e *colly.HTMLElement) {
		url := e.ChildAttr("a[data-event-action=title]", "href")
		if isImgur(url) {
			fmt.Println(url)
			browse.GoTo(ctxt, chrome, url)
			time.Sleep(duration * time.Second)
		}
		c.Visit(url)
	})

	// Access GFYcat video source
	c.OnHTML(".video-player-container", func(e *colly.HTMLElement) {
		videoSrc := e.ChildAttr("source", "src")
		fmt.Println(videoSrc)
		browse.GoTo(ctxt, chrome, videoSrc)
		time.Sleep(duration * time.Second)
	})

	// Continue to the next page on Subreddit
	c.OnHTML("span.next-button", func(e *colly.HTMLElement) {
		next := e.ChildAttr("a", "href")
		c.Visit(next)
	})

	// Enter subreddit as CLI argument eg. `./GoGiffy r/gifs 3`
	reddits := os.Args[1:]

	for _, reddit := range reddits {
		prefix := "https://old.reddit.com/"
		c.Visit(prefix + reddit)
	}
	c.Wait()
}

func getDuration() time.Duration {
	i, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	return time.Duration(i)
}

func isImgur(link string) bool {
	return strings.Contains(link, "i.imgur")
}
