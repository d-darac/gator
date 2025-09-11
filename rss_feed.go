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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error response:\n  status:  %s\n  message: %s", res.Status, string(data))
	}

	var rssFeed RSSFeed

	if err := xml.Unmarshal(data, &rssFeed); err != nil {
		return nil, fmt.Errorf("error unmarshaling xml: %w", err)
	}

	unescapeStringFields(&rssFeed)

	return &rssFeed, nil
}

func unescapeStringFields(feed *RSSFeed) {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Item[i] = item
	}
}
