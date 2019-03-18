package main

import (
	"context"
	"log"
	"os"

	"github.com/alexbotello/GoGiffy/browse"
	"github.com/alexbotello/GoGiffy/scraper"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Missing an argument: ./GoGiffy SUBREDDIT DURATION")
	}
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := browse.StartChrome(ctxt)
	scraper.Scrape(ctxt, c)
	browse.WaitWithChrome(c)
}
