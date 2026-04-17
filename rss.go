package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (f *RSSFeed) unescapeContent() {
	f.Channel.Title = html.UnescapeString(f.Channel.Title)
	f.Channel.Description = html.UnescapeString(f.Channel.Description)

	for i, rssItem := range f.Channel.Item {
		f.Channel.Item[i].Title = html.UnescapeString(rssItem.Title)
		f.Channel.Item[i].Description = html.UnescapeString(rssItem.Description)
	}
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	xml.Unmarshal(resBytes, &feed)

	feed.unescapeContent()

	return &feed, nil
}

func scrapeFeeds(s *state) {
	toFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	err = s.db.MarkFeedFetched(context.Background(), toFetch.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	feed, err := fetchFeed(context.Background(), toFetch.Url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		fmt.Printf("- %v\n", item.Title)
	}
}