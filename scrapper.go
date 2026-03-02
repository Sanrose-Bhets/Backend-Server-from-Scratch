package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Sanrose-Bhets/Backend-Server-from-Scratch/internal/database"
)

// this scrapping function works as long as the server is running,
// this function takes a connection to the database,
// no of concurrency = how many diff go routines you wanna do the fetch the diff feeds
// how much time inbetween request to scrape new RSS feeds
// a Long running job
func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	//created a no value for loops for te body of the ticker to immediately work and then wait, aka do while kinda thing
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feed:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapFeed(db, wg, feed)
		}
		wg.Wait()
	}

}

func scrapFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
		return
	}

	rssFeed, err := urlTOFeed(feed.Url)
	if err != nil {
		log.Println("Error Fetching Feed ", err)
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Found post", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))

}
