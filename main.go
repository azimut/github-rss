package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/gorilla/feeds"
)

func main() {
	login := os.Args[1]

	events, err := getEvents(login)
	if err != nil {
		log.Fatal(err)
	}

	feed := newFeed(login)
	for _, event := range events {
		item := parse(event)
		if item != nil {
			feed.Items = append(feed.Items, item)
		}
	}
	feed.Created = feed.Items[0].Created
	atom, err := feed.ToAtom()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(atom)
}

func newFeed(login string) *feeds.Feed {
	now := time.Now()
	feed := &feeds.Feed{
		Title:   fmt.Sprintf("%s github activity", login),
		Link:    &feeds.Link{Href: fmt.Sprintf("https://github.com/%s", login)},
		Created: now,
	}
	return feed
}

func getEvents(login string) ([]*github.Event, error) {
	client := github.NewClient(nil)
	events, _, err := client.Activity.ListEventsReceivedByUser(
		context.Background(), login, true, nil)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func parse(event *github.Event) *feeds.Item {
	switch *event.Type {
	case "WatchEvent":
		return feedWatch(event)
	case "CreateEvent":
		return feedCreate(event)
	}
	return nil
}

func feedWatch(e *github.Event) *feeds.Item {
	return &feeds.Item{
		Title: fmt.Sprintf("%s starred %s",
			*e.Actor.Login,
			*e.Repo.Name),
		Link: &feeds.Link{Href: fmt.Sprintf("https://github.com/%s",
			*e.Repo.Name)},
		Created: *e.CreatedAt,
	}
}

func feedCreate(e *github.Event) *feeds.Item {
	return &feeds.Item{
		Title: fmt.Sprintf("%s created %s",
			*e.Actor.Login,
			*e.Repo.Name),
		Link: &feeds.Link{Href: fmt.Sprintf("https://github.com/%s",
			*e.Repo.Name)},
		Created: *e.CreatedAt,
	}
}
