package browse

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

// StartChrome will start a new Chrome instance
func StartChrome(ctxt context.Context) *chromedp.CDP {
	// create chrome instance
	c, err := chromedp.New(ctxt)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// GoTo will navigate to given url
func GoTo(ctxt context.Context, c *chromedp.CDP, url string) {
	err := c.Run(ctxt, chromedp.Tasks{chromedp.Navigate(url)})
	if err != nil {
		log.Fatal(err)
	}
}

// WaitWithChrome will keep Chrome alive until shutdown
func WaitWithChrome(c *chromedp.CDP) {
	err := c.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
