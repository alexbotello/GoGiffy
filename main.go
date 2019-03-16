package main

import (
	"context"

	"github.com/alexbotello/GoGiffy/browse"
	"github.com/alexbotello/GoGiffy/scraper"
)

func main() {
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := browse.StartChrome(ctxt)
	scraper.Scrape(ctxt, c)
	browse.WaitWithChrome(c)
}
